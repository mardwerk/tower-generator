# Local character Wiki

A character entry is `wiki/<work>/<character>/README.md`, with retained source documents in `sources/*.md` beside it.
Folder names are readable Unicode slugs. The local Wiki is ignored by Git; the format, category definitions and tests are tracked.
Markdown and YAML are data. No MDX, executable page content, database or saved search index is required.

The entry's YAML frontmatter contains identity, canon scope, aliases, cited abilities and review metadata.
Its Markdown body contains the readable summary, evidence limits, contradictions and open questions.
Abilities, traits and equipment use the same record shape. Every record needs cited source passages and its applicable scope.

```markdown
---
formatVersion: 1
kind: character-wiki
name: Example Character
work: Example Manga
scope: Manga through chapter 20
language: en
aliases: []
status: draft
reviewedAt: null
sources: [supplied]
abilities:
  - id: distant-presence
    kind: ability
    name: Distant presence sensing
    description: Senses another person's presence at a distance.
    scope: Demonstrated in chapter 20
    limitations: The source does not establish continuous use or exact range.
    evidence:
      - source: sources/supplied.md
        passage: p-EXAMPLE
    classification:
      delivery: remote
      functions: [perception]
      status: suggested
---

# Example Character

A readable summary of the cited findings. Record unresolved questions here.
```

Use the actual passage ID shown by the Wiki in place of `p-EXAMPLE`.
A source file has YAML metadata with `id`, `title`, `url`, `retrievedAt` and `access`, followed by retained passages:

```markdown
---
id: supplied
title: Supplied evidence
url: null
retrievedAt: null
access: supplied
language: en
---

# Supplied evidence

<!-- passage:p-example -->
The retained passage, with its original reference when supplied manually.
```

Research entries and source documents use English. Collection requests English pages and rejects HTML pages explicitly declaring another language.
English reference pages may retain original-language character names; their research text and summaries remain English.
Retrieved sources require a public HTTP(S) URL and a timezone-aware retrieval date stored as text.
`resolvedUrl` records redirects; `excerpt: true` identifies an intentionally retained excerpt.
Supplied evidence must include enough origin information for human review; a passage without a URL is not automatically reliable.
New passage IDs derive from their content, so unrelated additions do not renumber citations. Keep IDs stable when editing source files by hand.

Collection accepts up to ten URLs and retains up to ten source documents per entry.
Saved URLs are reused without fetching unless refresh is explicitly selected. Repeated URLs and exact repeated passages are deduplicated.
Fetches and validation finish before file writes begin. Each file replacement is atomic; the entire directory is not a transactional database.
A failed fetch preserves the saved files. Collection preserves existing findings and summary unless a replacement summary is supplied.
If refresh removes a cited passage, retain that older passage with its original retrieval date in `retainedPassages` and reset the entry to draft.
The reader presents these passages as older evidence. Repeated claims across different sources retain their separate references.

One entry has one declared scope. Collection rejects a conflicting scope or another identity that normalizes to the same folder.
Use a distinct character/version name, such as `Usopp (anime)`, when separate entries are needed. Record narrower scopes on individual abilities.
Invalid entries appear as warnings in lookup; they do not hide valid entries. Paths outside the selected Wiki and child symbolic links are rejected.

[research-categories.yaml](research-categories.yaml) defines delivery and function categories independently of any Profile.
Classifications may have several functions, with `other` and `unknown` available. `suggested` tags are interpretations, not cited facts or game rules.
Set a classification to `reviewed` only after reviewing that classification. Manual review of an entry does not automatically approve its tags.

Saving an entry resets its review status to draft. Mark reviewed records an explicit human review of the cited findings and a content fingerprint.
Lookup detects later changes to entry or source files and presents the entry as a draft again. It performs no network request.
Concurrent editors send the opened revision; saving with a stale revision fails and preserves the files.
This is a bounded review record, not proof that a character catalog is complete or its sources are canon.

The reader returns a structured object in memory for the website and later consumers. Markdown files remain the only source of truth.
A consumer supplies the Wiki directory and character key. Profile mappings, upgrade choices, numerical stats and game adaptations remain outside this format.
