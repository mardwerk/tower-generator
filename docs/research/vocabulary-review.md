# Vocabulary review questions

These notes need review against local contracts. They do not establish current terms or accept behavior.

| Question | Proposed resolution to review |
| --- | --- |
| Tower Generator currently selects a single file (`unit.Profile`) for generation. Its `definition.profile` field remains in the mechanics Definition, and `docs/MECHANICS.md` calls that field "a Profile" once. | The accepted Profile is a directory selected by path (Planning audit #69). The single-file implementation awaits migration (tower-generator#85). The final name and file placement for `definition.profile` are deferred. |
| Tower Generator's design policy and finding messages say "manual ability"; Manga Mayhem's glossary says Active Ability. | Active Ability is the term. Tower Generator's `CONTEXT.md` maps it to the existing field names; the fields keep their names. |
| Result, Run, Candidate, Draft and Finding mean different things in Planning and in Tower Generator. | Separate scoped rows; name the scope outside the owning repository. |
| The td-profile contract defines a generic `validator` CLI. Tower Generator's current artifacts and Default Profile have not adopted this checker. | Pending compatibility and adoption work, separate from #100. This vocabulary proposal preserves the Reviewed Profile, Profile Template, Profile Validator and Tower Generator Default Profile meanings. It does not adopt the checker for Tower Generator or Manga Mayhem. |
