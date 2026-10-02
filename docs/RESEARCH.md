# Character research

The local website focuses on collecting, inspecting and manually reviewing character research.
It uses the retained Lab design system, centered navigation, one visible heading per page and no subheadings.
The plus opens Create; Create and Wiki are active; Library, Generations and the right-hand Settings control remain visible and disabled. Tower building and Profile controls are outside the website.
The research implementation is recorded in [#102](https://github.com/mardwerk/tower-generator/issues/102), and the compact Create layout in [#104](https://github.com/mardwerk/tower-generator/issues/104).

Install Python 3.10 or newer, Node.js 22.12 or newer, and the research dependencies:

```sh
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt
pnpm install
pnpm dev
```

Research does not require a checker installation, Go or an Atlas checkout.
`pnpm preview` uses the same local commands after `pnpm build`. Existing Tower experiment files remain available separately through [Authoring](AUTHORING.md).

Create is the starting page. Its base view shows character and anime/manga inputs, Research, and small Research inputs and Wiki controls.
Research inputs opens the canon scope, up to ten English source URLs, refresh and optional supplied passages and summary.
Submitting without scope or evidence opens that panel with a short message. The form retains unsent inputs when switching tabs.
Collect more evidence opens an existing entry's inputs for editing; evidence, classifications and review stay in Wiki.
Collection extracts readable page text with Trafilatura and retains the original URL, retrieval date and source passages. It requests English pages and rejects HTML explicitly declaring another language.
Saved URLs are reused by default; Fetch saved URLs again requests a refresh. A failed batch leaves previously saved evidence intact.
Collection creates a draft in `wiki/<work>/<character>/`. It does not automatically search, synthesize canon facts or call Jev.

Open an entry from Wiki and select Edit entry to add sourced abilities, traits, equipment, limitations and open questions.
Use the source and passage IDs displayed below the entry. The [Wiki format](wiki-format.md) contains a complete record example.
Saving validates references and categories before replacing the Markdown file. Mark reviewed records manual review after cited records have been added.
Tag review is separate from entry review. Suggested categories remain suggestions until their classification status is explicitly changed.
Export Markdown downloads the entry; copy its character directory to retain the referenced source files for another consumer.

The CLI owns storage, retrieval, validation, deduplication, refresh and review. The website owns forms, selection and presentation.
The local Vite server invokes `python3 -B -m src.wiki --json` with structured input and no shell. It contains no separate persistence implementation.
The same commands work without a browser; use `pnpm wiki --help` or `python3 -B -m src.wiki --help`:

```sh
pnpm wiki list --query Usopp
pnpm wiki show one-piece/usopp
pnpm wiki collect --name Usopp --work "One Piece" --scope "Manga through Dressrosa" \
  --url https://en.wikipedia.org/wiki/Usopp
pnpm wiki collect --name "Example Character" --work "Example Manga" --scope "Chapter 20" \
  --evidence "Referenced passage supplied by the researcher."
pnpm wiki save example-manga/example-character --file /path/to/edited-README.md
pnpm wiki review example-manga/example-character
pnpm wiki categories
```

Use the exact saved scope when adding or refreshing evidence. Add `--refresh` to explicitly fetch saved URLs again.
`--wiki /path/to/wiki` selects another directory and `--json` returns structured results; put these options before the command.
`collect --input request.json` and `save <key> --input request.json` accept the same JSON requests as the website. Use `--input -` for standard input.
Save requests contain `markdown` and optionally the opened `revision`. Review also accepts `--expected-revision`.
The program keeps no session, database or history outside the selected files; lookup builds its index in memory.

Migrate earlier local JSON sources explicitly:

```sh
pnpm wiki migrate
```

Migration copies valid records from `sources/*.json` into the Wiki as drafts, retains their evidence and skips existing character entries.
It leaves original files intact. Existing scopes and manually authored Wiki records are not overwritten.
The tracked [Usopp JSON](../sources/usopp.json) remains a migration example and the earlier Tower experiment's source snapshot input.

Start with agent-assisted source discovery and synthesis for Usopp, then Luffy and Gojo. Review the cited findings and tags manually.
Measure missing facts, source failures, mistaken merges, tag corrections and review time before expanding automation.
The next automation work in [#103](https://github.com/mardwerk/tower-generator/issues/103) should search for identity and abilities, select up to ten relevant distinct pages, synthesize cited records, and ask a decision model about those records.
Jev should receive compact facts, relevant passages, the [category definitions](research-categories.yaml) and explicit unknown outcomes.
Web ranking alone does not determine source quality; source limits and manga/anime/version differences remain visible.
This current iteration adds no paid model call or provider configuration.

`pnpm test` covers Wiki round trips, CLI migration and offline lookup, evidence reuse, failed batches, stable passages, refreshed citations, review changes and invalid references.
Run `pnpm typecheck`, `pnpm format:check` and `pnpm build` for the website. Browser verification covers evidence collection, saving, review and responsive layout.
