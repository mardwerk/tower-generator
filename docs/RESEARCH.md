# Character research

Run name-only research from the CLI or the single-field Create page:

```sh
pnpm wiki research usopp
pnpm wiki research Sogeking
pnpm wiki research "Satoru Gojo"
pnpm wiki research "Sakura" --series Naruto
```

Research searches English references, resolves canonical identity, checks local names and aliases, fetches evidence, and saves cited draft findings.
A repeat verifies and expands the existing entry with the same allocation. It keeps the entry's canon scope, evidence and manual prose.
The implementation and reviewed scope are recorded in [#103](https://github.com/mardwerk/tower-generator/issues/103).

The CLI requires Python 3.10 or newer and the research dependencies. Node.js 22.12 or newer is needed for website development.
`python3 -B -m src.wiki research usopp` runs directly; the `pnpm wiki` examples below are convenience shortcuts.
For the website development environment:

```sh
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt
pnpm install
pnpm dev
```

Set `OPENROUTER_API_KEY` and `OPENROUTER_MODEL` in the process environment or the ignored local `.env`.
Environment values override the file. `OPENROUTER_REASONING` is optional and defaults to low.
The key stays in the protected configuration and is sent only to OpenRouter; it never enters Wiki files or prompts.
Research uses OpenRouter's Exa web plugin and the selected model, with no automatic model fallback or retries.
It does not require the checker, Go or an Atlas checkout.

Each run allocates at most 10 search results, 10 page attempts and 2 model calls.
The calls allow 2,500 and 8,000 output tokens; each model input is limited to 100,000 characters including its instructions.
Use `--pages 1` through `--pages 10` to lower the search/page allocation. Fresh and repeat runs use identical configured limits, not identical actual cost.
The last run records limits, actual usage/cost reported by the provider, attempted sources, failures, model and category/normalization versions.
Available sources may consume less than the maximum; a repeat can improve verification without finding a new capability.

Existing entries guide search through their inventory, source gaps and follow-up questions.
Up to three older URLs are rechecked, with the remaining page allocation used for searched references. Up to 20 least-recently checked records are verified per run.
New cited findings are added; unreviewed drafts can be refined. Source passages cited by findings, verification results and proposed updates survive refresh with their old retrieval date.
Human-reviewed findings and classifications remain protected across later runs. Proposed changes to those records are kept in `suggestedUpdates` for manual review.
Model-supported, conflicting and unresolved checks are stored on records as `verification`. These are model assessments, not human acceptance or proof of canon.
Unsupported findings are not deleted merely because today's sources omit them. New classifications remain suggested; reviewing an entry does not approve its tags.

Identity matching first compares the resolved series and normalized canonical names/aliases against entry metadata.
Unicode normalization, case folding, accent folding and punctuation removal handle older spellings for matching. Folder naming keeps the existing Unicode lowercase/hyphen rules.
Research scans older folder spellings and normalizes a confirmed matching directory while preserving all its files. It never merges two existing character folders silently.
A 99% text match needs a shared identity source URL before reuse; otherwise it is reported as a possible duplicate. Similarity is not an identity probability.
Ambiguous names require a series hint. Add a confirmed alias to an existing entry when a likely duplicate needs explicit resolution.
`--scope "Manga through Dressrosa"` sets a new entry's canon boundary; a conflicting scope on an existing entry is rejected.

Direct fetching requests English pages and rejects explicitly non-English HTML. Collection extracts readable text with Trafilatura.
When a page is blocked, a usable English search excerpt may be retained with `retrievalMethod: search-excerpt` and `excerpt: true`.
Search excerpts can be incomplete or stale; source failures and limits remain visible. Predominantly Japanese/Chinese excerpts are rejected; untagged or supplied text still needs manual language review.
Search-selected references are filtered against the returned citations; unsupported URL suggestions are ignored and reported.
Generated role-play biographies from the DaddyJim site observed during the live test are excluded from model evidence, while any saved originals remain on disk.
Only source URLs returned by search are accepted for identity evidence, and generated findings must cite passages supplied to the model.
All fetching, synthesis and validation finish in a temporary directory before replacing the live character directory. Failure or a detected stale edit preserves the existing entry.
Directory replacement has a short rename window; this is a local tool, not a multi-user database. The browser serializes mutations and CLI callers should avoid overlapping writes.

The website uses the retained Lab styling, centered navigation, one heading per page and no subheadings.
Create shows a character input and Research. Series/scope hints and manual references stay behind Research inputs.
Without supplied URLs/passages, Research invokes name-only research. Supplied references use the existing manual collection command and require series and scope.
Wiki presents entries, evidence and manual editing/review. The CLI owns persistence and validation; the local server invokes that same CLI with structured input and no shell.
`pnpm preview` uses the same commands after `pnpm build`. The earlier Tower experiment remains separate in [Authoring](AUTHORING.md).

Inspect and manually edit the portable files with the CLI or Wiki:

```sh
pnpm wiki list --query Usopp
pnpm wiki show one-piece/usopp
pnpm wiki collect --name Usopp --work "One Piece" --scope "Manga through Dressrosa" \
  --url https://en.wikipedia.org/wiki/Usopp
pnpm wiki save one-piece/usopp --file /path/to/edited-README.md
pnpm wiki review one-piece/usopp
pnpm wiki categories
pnpm wiki migrate
```

Manual collection reuses saved URLs unless `--refresh` is selected. It does not synthesize findings or automatically search.
The exact saved scope is required for manual additions. Saving validates references and categories; review records human acceptance of findings without automatically approving tags.
`research --input request.json` accepts `name`, optional `series`, `scope` and `pages`; `--input -` reads standard input, as used by the website.
`collect --input request.json` and `save <key> --input request.json` accept their existing structured requests.
Put `--wiki /path/to/wiki` and `--json` before the command to select another directory and request machine-readable output.
Lookup is offline and rebuilds its index from local files. There is no hidden session or database.

The [Wiki format](wiki-format.md) describes identity, findings, citations and review metadata.
Migration copies earlier `sources/*.json` into draft Markdown without deleting originals or overwriting existing entries.
The tracked [Usopp source](../sources/usopp.json) remains a migration example and the parked Tower experiment's snapshot input.
Research stays independent of any Profile, Tower upgrade structure or numerical balance. Suggested categories use [research-categories.yaml](research-categories.yaml).
The separate [Jev comparison](https://github.com/mardwerk/tower-generator/issues/105) can use manually reviewed findings and tags as its baseline; this pipeline uses the already configured model.

Run `pnpm test`, `pnpm typecheck`, `pnpm format:check` and `pnpm build` for the focused checks.
Tests cover equal-budget repeat research, normalization/aliases/old folders, protected reviews, failed sources, bad citations, ambiguous identities and stale edits, plus the existing storage/authoring checks.
