# Triage compatibility

The repository's [label rules](../ISSUE_LABELS.md) govern work state and closure.
Upstream triage roles describe questions or readiness; they do not replace this lifecycle.
Apply these compatibility rules instead of upstream's exactly-one-category and exactly-one-triage-state defaults.

| Upstream role | Repository handling |
| --- | --- |
| needs-triage | Use `status:triage` for unresolved scope, ownership or scheduling without recorded selection. |
| needs-info | Record the missing information and who can supply it in the issue; preserve its recorded work state. |
| ready-for-agent | Record the accepted outcome, reviewed design contract, prerequisites and verification evidence; preserve its recorded work state. |
| ready-for-human | Record the action that requires a person and the reason; preserve its recorded work state. Use `review:decision` only for an actual pending owner decision. |
| wontfix | Distinguish completed work, duplication, invalid reports and explicit rejection. Apply native closure and the applicable existing closure rules. Retain deferred or excluded proposals when the owner selected that state. |

Discover triage work through `status:triage` and relevant existing issues, then read their bodies and comments for missing information and readiness.
Absence of upstream role labels does not make an issue untriaged; preserve owner-selected states.

Keep rejection and exclusion evidence in the owning issue and established documents instead of creating an `.out-of-scope/` decision store.
Keep `status:out-of-scope` proposals open unless an explicit owner decision changes their disposition.

Use `type:bug`, `type:feature`, `type:documentation` or `type:maintenance` when appropriate; type is optional.
Do not introduce upstream default labels or map implementation readiness to `status:active`.
Only a recorded owner decision selects, defers, pauses or rejects work.
Readiness requires the current accepted design contract; an active issue can still await a decision or prerequisite.
