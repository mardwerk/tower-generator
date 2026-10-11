# Domain conventions

This repository has one product context.
Two delivery modes do not by themselves require multiple glossaries or a monorepo layout.

Before exploration, read [README](../../README.md), [product rules](../PRODUCT.md) and the relevant owning document.
Read [research rules](../RESEARCH.md) for identity, evidence and confidence, and [Profile ownership](../PROFILE.md) for game constraints and attribution.
Read applicable recorded issue decisions and local ADRs when present.

Agents with Planning access also read its [master glossary](https://github.com/mardwerk/planning/blob/main/GLOSSARY.md).
Only Reviewed rows establish controlled terms; pending rows and pinned retired CONTEXT documents are proposals or historical evidence.
Public contributors use the product contracts without requiring private Planning access.
Keep the controlled vocabulary in Planning; do not create a second local controlled term table.
Propose new controlled terms or changed meanings through Planning's existing vocabulary workflow.

Product behavior and acceptance remain in their owning documents and issues.
Use `docs/adr/` lazily for durable, costly-to-reverse choices with meaningful alternatives and non-obvious reasons.
Link existing decisions rather than copying their contracts into new ADRs.
Flag conflicts between proposed changes and an accepted contract before implementing them.
