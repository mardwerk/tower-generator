# Web app (UnitLab)

The web app is a React client served by `mardwerk-unit serve`, which embeds it. The client talks only to the local `/api/v1` API; every rule, prompt and check runs in the Go server ([ARCHITECTURE.md](ARCHITECTURE.md#http-api)).

## Start

```sh
go build -o mardwerk-unit ./src/cli
./mardwerk-unit serve
```

Open `http://127.0.0.1:4317`. The terminal also shows the model and the OpenRouter key in use, the key masked, and the `.env` file or variable they came from. Use `--port 4318` for another port and `--provider codex` to use an existing Codex login instead of OpenRouter. Stopping the server (Ctrl+C) cancels running generations.

When you change the client, rebuild it with `pnpm build` (Node.js 22 or newer) and rebuild the binary; `src/web/dist` is what the binary serves.

## Generate a unit

Enter a character name and select Generate. The Profile dropdown below the name, beside Inputs and rules and Import, picks the rules; the BTD6-inspired default is preselected. The server looks the character up on Wikipedia (with Wikidata and Fandom for text and images) and asks you to choose when the name is ambiguous. It saves the found Sources to your library, prepares them under the Profile, then drafts, checks and reviews the unit. Missing sources produce an error, never invented canon.

- Starting a generation clears the name field and hides earlier failed or stopped runs from the activity bar; their revisions stay available.
- A generation that fails its checks after the repair attempts shows **Download failure evidence** below the error: the rejected plans and outputs with their issues, the accepted plan if there was one and the source passages the model was given, so the failure can be read without generating again.
- Ctrl-click or Cmd-click Generate to stay on the create page; several runs can proceed at once, each with its own Stop.
- **Inputs and rules** lets you edit or import a request before generating. An imported request keeps its own rules until you pick a Profile; otherwise the selected Profile's rules apply when it is prepared.
- The sidebar runs each stage (prepare, draft, check, review) separately. Rerunning a stage creates a new revision and keeps the old one. Continue runs the remaining stages; Stop cancels the current one.
- The unit sheet shows references on the left, the unit in the middle, and stages and usage on the right. It opens with the `0-0-0` description (price and attack). Each path is labeled top, middle or bottom and each purchase by build code (`1-x-x`, `x-4-x`, `x-x-5`); a purchase that adapts a named technique shows the plan's description of that adaptation, which the review checks against its effects; opening a purchase shows its exact effects. **Crosspaths** shows one grid per pair of paths, first path down and second across, with the total price of every legal build (12 early and 36 advanced under the default); selecting a build shows what each path adds and the resulting attack, and **All crosspath builds** lists them as tables. A revision gets **Patch notes** with changed mechanics apart from renamed purchases. Checks, findings (including unsupported mechanics), reserved techniques and open decisions are in the panels below the sheet, apart from the unit. When today's authoring checks flag a unit whose sheet still resolves, such as a Result saved before a check was added, a notice under the header says so, and the checks panel lists them as current checks, apart from the stored findings.
- Feedback creates a revision that keeps confirmed decisions.
- Closing or reloading the tab cancels running generations and loses unsaved edits.

## Profiles

A Profile is the set of rules a unit is generated under: a mechanics Definition, its rules text and the task. The **Profiles** tab lists the built-in, read-only Profiles and the Profiles saved in the Profiles folder (`data/profiles`, `--profiles DIR`). The BTD6-inspired default has three paths of five tiers, BTD6 crosspath rules, 150 starting health, the character design rules and scale references for several roles (group capacity, precision, attack speed, actives, control and support) pinned to btd6-atlas capture 56.3. The built-in stacking example adds poison and bleed, which stack up to a limit and a combined cap. The tab shows the selected Profile's paths and tiers, prices, stats and limits, and for a version 2 Definition its status effects (kind, bounds, duration limit, stacking and immunities), damage types, targeting, detection, enemy properties and the properties that accept bonus damage. **Duplicate** makes an editable copy; saving runs the same checks as preparing a request, so a numerical Profile must keep three paths of five tiers. To add a status effect, damage type, targeting mode or detection trait, edit the Definition's `vocabulary` ([contract](MECHANICS.md#profile-defined-vocabulary-version-2)); the copy's version follows its Definition. **Use for new units** selects it on the Generate form. A unit's kit shows its status effects, detection and bonus damage (as +N against a property) with the Profile's names.

Profiles saved by earlier versions lived in `data/runs/library/profiles`; move them to `data/profiles` to keep using them.

## Settings

- **Provider.** OpenRouter is the default and uses `openrouter/free` (free models only, no paid fallback). A key is required even for free models: set `OPENROUTER_API_KEY` in `.env` or the environment, or enter it in Settings, where it stays in server memory and never enters an artifact. The top bar shows the key in use, masked (for example `sk-or-v1-abc...xyz`), with the model in its tooltip; Settings adds where the key came from: `.env`, the environment, or Settings. A configured key is not necessarily one OpenRouter accepts. `OPENROUTER_MODEL` and `OPENROUTER_REASONING` in `.env` set OpenRouter defaults. `CODEX_MODEL` and `CODEX_REASONING` set local Codex defaults. Settings fetches the selected provider's model catalog for a dropdown and shows that model's reasoning levels as choices. Saving changes the local server setting until it stops.
- **Library folder.** Defaults to `data/runs/library`; a folder chosen in Settings is recorded in `data/runs/lab-settings.json`. Both are ignored by Git.

## Library

Completed Results and researched Sources are saved automatically; earlier stages can be saved with Save to library, and the status line then names the file it wrote. Each character has a folder under its work, both named readably:

```text
data/runs/library/
  bloons-td-6/                                  the work (anime, game or other source)
    dart-monkey/                                the character
      character.json                            the exact name and work this folder belongs to
      dart-monkey.sources.1a2b3c4d5e6f.json     researched Sources
      dart-monkey.result.0eb43da91c2f.json      a reviewed unit (one file per stage and revision)
      dart-monkey.result.0eb43da91c2f.md        its Markdown render
      assets/                                   icons, image receipts and the portrait choice
```

The 12 hex digits are the start of the record's SHA-256 identity, so stages and revisions never share a file, and the library still recognizes a record by that identity, not by its name. Names are reduced to lowercase letters and digits joined by hyphens; a character whose name reduces to a folder that another character already owns gets `<name>-<hash>` instead, so similar names never share a folder. Library cards show each record's path. The **Units** tab lists generated units grouped by work with portraits; the **Research** tab lists saved Sources, and opening one prepares it under the selected Profile without researching again. Generate does the same for a name: when saved Sources were researched under that name or describe a character of that name, it reuses the newest ones instead of searching online, and asks which character when several match. **Find references again** in the generation flow searches online anyway and extends the saved Sources: a page retrieved again replaces its older copy, pages not found again stay, and identical documents are not duplicated. When that search asks which character, the chosen page extends the same saved Sources. The result is saved as new Sources. Clean up and Clear act on the open tab and remove only those record files and their renders; character folders, markers and icons stay.

Records saved by earlier versions (`unitlab-<id>.json` at the top of the library, icons under `assets/unit-<hash>`) stay readable and are listed with their old path. **Arrange by source and character** (or `mardwerk-unit library migrate`) moves them into their folders, keeping each record's identity and saved time; a file is removed only after its copy reads back intact, and anything it cannot place stays where it is.

Below each unit sheet, Save to library stores it, and the Export menu downloads it as JSON (compatible with the CLI), Markdown, or a session file that keeps unsaved editor content. With more than one revision, the revision and comparison selectors appear above them.

## Icons and portraits

Every attack, upgrade and ability has an icon placeholder with a copyable image prompt (and a Codex prompt that saves the PNG) and a fixed PNG destination in the character's `assets` folder. You can also generate one icon through OpenRouter after an explicit confirmation that names the model, the estimated price and the destination. The default image model is `meta/muse-image` (about $0.01 per image; `OPENROUTER_IMAGE_MODEL` changes it). There are no retries or fallbacks. Results are validated and converted to PNG (8 MB limit), and the reported cost is kept beside the image. A gallery image can be chosen as the unit's portrait; that preference never changes the unit or its hash.

## Security

The server binds `127.0.0.1`, checks the host and origin of every request, and requires a fresh session token that the page receives automatically. It writes only inside the library and Profiles folders and to `data/runs/lab-settings.json`.
