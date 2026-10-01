"""Check the Default Profile against Atlas with a released or local checker."""

import argparse
import hashlib
import json
from pathlib import Path
import platform
import subprocess
import sys


PROJECT = Path(__file__).resolve().parents[1]
DEFAULT = PROJECT / "default-profile"
CAPTURE = "data/56.3-build-24829026/game-data"
DART = "Towers/DartMonkey/DartMonkey.json"


def check_schemas(expected, profiles):
    shared = {p.relative_to(expected): p.read_bytes() for p in expected.rglob("*.json")}
    if not shared:
        raise ValueError(f"No shared schemas found in {expected}")
    for profile in profiles:
        actual = {
            p.relative_to(profile / "schemas"): p.read_bytes()
            for p in (profile / "schemas").rglob("*.json")
        }
        differences = sorted(str(p) for p in shared.keys() | actual.keys()
                             if shared.get(p) != actual.get(p))
        if differences:
            raise ValueError(f"Schema drift in {profile}: {', '.join(differences)}")
    print(f"Shared schemas match: {len(shared)} files", flush=True)


def checker(args):
    if args.local_checker:
        destination = PROJECT / ".tools/td-profile-local"
        destination.mkdir(parents=True, exist_ok=True)
        binary = destination / ("validator.exe" if sys.platform == "win32" else "validator")
        print(f"Building local checker from {args.td_profile}", flush=True)
        subprocess.run(["go", "build", "-o", str(binary), "./cmd/validator"],
                       cwd=args.td_profile, check=True)
        return binary, args.td_profile / "profile/schemas"

    pin = json.loads((DEFAULT / "validator-release.json").read_text())
    architecture = {"x86_64": "amd64", "aarch64": "arm64"}.get(
        platform.machine().lower(), platform.machine().lower())
    if (platform.system().lower(), architecture) != (pin["goos"], pin["goarch"]):
        raise ValueError("The pinned binary targets Linux amd64; use --local-checker on this platform.")
    binary = DEFAULT / "validator"
    if not binary.exists():
        binary = args.atlas / "profile/validator"
    if hashlib.sha256(binary.read_bytes()).hexdigest() != pin["executableSha256"]:
        raise ValueError(f"Pinned checker checksum mismatch: {binary}")
    atlas_pin = json.loads((args.atlas / "profile/validator-release.json").read_text())
    if atlas_pin["release"] != pin["release"]:
        raise ValueError("Atlas uses a different checker release; review the pins before adopting it.")
    print(f"Using released checker {pin['release']}: {binary}", flush=True)
    return binary, args.atlas / "profile/schemas"


def score(binary, profile, data, tower, label):
    result = subprocess.run(
        [str(binary), "score-tower", "--profile", str(profile),
         "--game-data", str(data), "--tower", tower],
        capture_output=True, text=True,
    )
    if not result.stdout:
        raise ValueError(result.stderr.strip() or f"{label}: checker produced no report")
    report = json.loads(result.stdout)
    points = report.get("score", {}).get("points")
    print(f"{label}: score={points}/100; files={report['filesChecked']}; errors={len(report['errors'])}")
    identity = report.get("profile", {})
    print(f"  Profile: {identity.get('id')} {identity.get('revision')}; sha256={identity.get('sha256')}")
    identity = report["checker"]
    print(f"  Checker: {identity['version']}; revision={identity.get('revision')}; modified={identity['modified']}")
    for error in report["errors"]:
        print(f"  {error['file']}#{error['pointer']}: {error['code']}: {error['message']}")
    if result.returncode or points != 100 or not report["score"]["complete"]:
        raise ValueError(f"{label}: expected complete validation and a score of 100")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--local-checker", action="store_true", help="rebuild the sibling td-profile checker")
    parser.add_argument("--atlas", type=Path, default=PROJECT.parent / "btd6-atlas")
    parser.add_argument("--td-profile", type=Path, default=PROJECT.parent / "td-profile")
    parser.add_argument("--game-data", type=Path, help="also check this candidate data directory")
    parser.add_argument("--tower", help="candidate Tower path relative to its game-data")
    args = parser.parse_args()
    if bool(args.game_data) != bool(args.tower):
        parser.error("--game-data and --tower must be supplied together")
    args.atlas = args.atlas.resolve()
    args.td_profile = args.td_profile.resolve()
    binary, schemas = checker(args)
    check_schemas(schemas, [args.atlas / "profile", DEFAULT])
    # ponytail: this baseline assumes Atlas bindings; revisit when the Default format changes.
    score(binary, args.atlas / "profile", args.atlas / CAPTURE, DART, "Atlas Dart")
    score(binary, DEFAULT, args.atlas / CAPTURE, DART, "Default Dart")
    if args.game_data:
        score(binary, DEFAULT, args.game_data.resolve(), args.tower, "Candidate")


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, subprocess.CalledProcessError) as error:
        sys.exit(f"check-default-profile: {error}")
