"""Build the latest released validator for this machine and prepare the example."""

import argparse
import hashlib
import json
import os
from pathlib import Path
import platform
import shutil
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True

from src.author import PROJECT, check, default_atlas, generate, write_json


DEFAULT = PROJECT / "default"
UPSTREAM = "mardwerk/td-profile"


def install_validator():
    release = subprocess.check_output(
        ["gh", "release", "view", "--repo", UPSTREAM, "--json", "tagName"], text=True,
    )
    tag = json.loads(release)["tagName"]
    goos = {"Linux": "linux", "Darwin": "darwin", "Windows": "windows"}.get(platform.system())
    goarch = {"x86_64": "amd64", "amd64": "amd64", "aarch64": "arm64", "arm64": "arm64"}.get(platform.machine().lower())
    if not goos or not goarch:
        raise ValueError(f"Unsupported native target: {platform.system()} {platform.machine()}")
    binary = DEFAULT / ("profile-validator.exe" if goos == "windows" else "profile-validator")
    metadata = DEFAULT / "validator-release.json"
    if binary.exists() and metadata.exists():
        identity = json.loads(metadata.read_text())
        if (identity["release"], identity["goos"], identity["goarch"]) == (tag, goos, goarch):
            if hashlib.sha256(binary.read_bytes()).hexdigest() == identity["executableSha256"]:
                print(f"Validator {tag} already built for {goos}/{goarch}", flush=True)
                return
    with tempfile.TemporaryDirectory(prefix="tower-validator-") as temporary:
        source = Path(temporary) / "td-profile"
        subprocess.run(["git", "-c", "advice.detachedHead=false", "clone", "--quiet", "--depth", "1", "--branch", tag,
                        f"https://github.com/{UPSTREAM}.git", str(source)], check=True)
        expected = {p.relative_to(source / "profile/schemas"): p.read_bytes()
                    for p in (source / "profile/schemas").rglob("*.json")}
        actual = {p.relative_to(DEFAULT / "profile/schemas"): p.read_bytes()
                  for p in (DEFAULT / "profile/schemas").rglob("*.json")}
        if not expected or expected != actual:
            raise ValueError("Latest release schemas differ from the Default Profile; track the update before adopting it")
        staged = Path(temporary) / binary.name
        environment = {**os.environ, "GOOS": goos, "GOARCH": goarch, "CGO_ENABLED": "0"}
        print(f"Building td-profile {tag} for {goos}/{goarch}", flush=True)
        subprocess.run(["go", "build", "-trimpath", "-o", str(staged), "./cmd/validator"],
                       cwd=source, env=environment, check=True)
        empty = Path(temporary) / "empty"
        empty.mkdir()
        result = subprocess.run([str(staged), "--profile", str(DEFAULT / "profile"),
                                 "--game-data", str(empty)], capture_output=True, text=True, check=True)
        report = json.loads(result.stdout)
        if not report["valid"] or report["filesChecked"] != 0:
            raise ValueError("Built validator did not pass the empty-data check")
        identity = {"repository": f"https://github.com/{UPSTREAM}", "release": tag,
                    "sourceCommit": subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=source, text=True).strip(),
                    "goos": goos, "goarch": goarch, "checker": report["checker"],
                    "executableSha256": hashlib.sha256(staged.read_bytes()).hexdigest()}
        shutil.copy2(staged, binary)
        other = DEFAULT / ("profile-validator" if goos == "windows" else "profile-validator.exe")
        other.unlink(missing_ok=True)
        notices = DEFAULT / "licenses"
        notices.mkdir(exist_ok=True)
        for name in ["LICENSE", "NOTICE.md"]:
            shutil.copy2(source / name, notices / name)
        shutil.copytree(source / "licenses", notices / "dependencies", dirs_exist_ok=True)
        write_json(metadata, identity)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--atlas", type=Path, default=default_atlas())
    args = parser.parse_args()
    install_validator()
    design = json.loads((DEFAULT / "usopp.json").read_text())
    generate(design, args.atlas.resolve(), DEFAULT)
    check(DEFAULT, design["id"])
    print("Ready. See docs/AUTHORING.md to inspect the Usopp Tower experiment.")


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, KeyError, StopIteration, subprocess.CalledProcessError) as error:
        sys.exit(f"setup: {error}")
