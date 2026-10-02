# Character research

The Lab has Create, Library and Tower views using the retained pre-cleanup design system.
Each page has one visible heading and no subheadings. Create collects character evidence; Library opens it for reuse; Tower inspects and edits the Usopp experiment.

Enter a character, work and canon scope, then provide source URLs or paste evidence under Supplied evidence and research notes.
Research and save retrieves the supplied public HTML or plain-text pages, retains their text and metadata, and saves a character source.
This is evidence collection, rather than an automatic judgment that a page matches the character or proves canon. Review the identity, scope, passages and source quality.
Research notes can summarize findings with document IDs, as the Usopp sample does; keep prices, Profile rules and Tower adaptations in the Tower design.

Library opens saved sources without fetching the web again. Export source downloads the portable JSON; the top-bar Import control validates and saves that same format.
Update research opens the existing character and source URLs in Create. Nothing is fetched until Research and save is selected.
Refresh replaces requested pages and retains other evidence documents. A failed fetch leaves the previously saved source unchanged.
Source-page size limits fail visibly instead of silently cropping evidence; a focused supplied excerpt can be used when a page is inaccessible or too large.

Sources live in `sources/<id>.json`. The curated [Usopp source](../sources/usopp.json) is tracked; newly collected files stay local and ignored by Git.
A character with the same name, work and scope reuses its source ID. A different scope gets another source, so research periods stay distinguishable.
Files are UTF-8 JSON with `formatVersion: 1` and `kind: character-source`:

```json
{
  "formatVersion": 1,
  "kind": "character-source",
  "id": "example-character",
  "character": {
    "name": "Character name",
    "work": "Source work",
    "scope": "Selected canon period"
  },
  "documents": [
    {
      "id": "reference-1",
      "title": "Source title",
      "url": "https://example.com/character",
      "retrievedAt": "2026-10-02T10:00:00+00:00",
      "access": "retrieved",
      "text": "Retained source passage."
    }
  ],
  "notes": "Findings and evidence limits, citing reference-1."
}
```

Retrieved documents need a URL and timezone-aware retrieval date. Supplied documents use `access: supplied` and can use null for those fields.
`resolvedUrl` records the final page after redirects; `excerpt: true` identifies a deliberately retained excerpt.
Source documents need unique IDs. The loader checks the format, evidence, dates, URLs, size limits and safe source identifiers without accessing the network.
This format is Tower Generator's small consumer interface. It does not replace the shared research workstream's future dossier contract.

Use for Usopp opens the matching source in the Tower workspace. Edit design changes the editable JSON; Build and check uses the installed validator and the sibling Atlas checkout.
The complete candidate is generated and checked in a temporary directory before the saved design or game-data changes. Invalid inputs preserve the last saved draft.
General character generation remains later work; sources for other characters can already be collected, inspected and exported.

The design's `sourceId` selects a saved character source. Generation retains its exact JSON in `default/character-source.json` and its SHA-256 identity in `default/source.json`.
Research updates do not change that snapshot until another build. The Profile and game-data stay separate from character evidence.
A saved source is reusable independently of the Profile; another consumer can read the exported JSON without a Tower Generator installation.

The Lab uses the existing local Vite middleware and Python standard library. It writes only source files and the fixed Usopp workspace.
No database, provider settings, paid model call or old generation engine is required.
`python3 setup.py` installs the checker once; `pnpm dev` starts the Lab. `pnpm preview` supports the same local operations after `pnpm build`.

Run `pnpm test` for source round trips, offline reuse, refresh and identity boundaries, plus authoring and source-snapshot checks.
Browser checks cover the Library flow, evidence collection, design persistence across views, positive and failed builds, and one heading on desktop and mobile.
Work and the reviewed implementation plan are recorded in [#101](https://github.com/mardwerk/tower-generator/issues/101).
