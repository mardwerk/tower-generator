# Issue labels, milestones and views

These rules apply to this repository only.

## Label rules

State, topic, type, review, origin and importance describe separate questions. Applying a label does not accept a contract or select work for the owner.

- Use one recorded state per open issue. Normal completion uses native closure without a terminal label. Remove the open state on closure; add at most one terminal reason when the recorded disposition supports it.
- Use topics for the responsibilities involved. Multiple topics are allowed when they apply to the same task. Split independent tasks into separate issues.
- Type is optional, with at most one primary change type.
- Review, Origin and Importance are optional. Only applied labels need a recorded reason. An absent importance label means unassessed.
- Importance is priority, independent of state. High means high priority, medium means ordinary priority, and low means lower priority. Set it from an owner decision or owner-requested assessment. Record blocking dependencies separately in the issue.
- Use `review:decision` for a named pending owner decision, `review:audit` for a fresh audit and `review:closure` for a proposed closure. Multiple review labels require distinct pending tasks. A closure proposal does not close, reject or change the issue's state.
- Use origins only for explicit source or product relationships. They do not create implementation dependencies.
- A milestone records a repository outcome and observable completion criteria. GitHub permits at most one milestone per issue. Assignment does not select work or accept a contract; use linked dependencies for other outcomes.
- Retained excluded proposals stay in their owning repository. Review may apply to them, but ordinary open-issue and owner-review lists exclude `status:out-of-scope`.

## States

States record owner decisions.

| Label | Meaning |
| --- | --- |
| `status:active` | Owner-selected current work or decision. |
| `status:triage` | Unresolved scope, ownership or scheduling without recorded selection. |
| `status:deferred` | Retained agreed later work or a pending prerequisite. |
| `status:parked` | A future option awaiting owner selection. |
| `status:on-hold` | An explicit owner pause with a reason and resumption conditions. |
| `status:out-of-scope` | An explicitly excluded proposal retained open. |
| `status:invalid` | A closed incorrect or non-actionable report with a recorded reason. |
| `status:duplicate` | A closed duplicate with a recorded owning issue. |
| `status:declined` | A closed explicit owner rejection. |

## Topics

These topics describe local responsibilities.

| Topic | Meaning |
| --- | --- |
| `topic:repository` | Repository organization, contributor rules, tracking and governance |
| `topic:operations` | Builds, test environments, publishing, deployment and operational data |
| `topic:interface` | User controls, presentation and interaction |
| `topic:api` | The Go `serve` API: REST operations, SSE progress, research queue, Wiki storage, startup and configuration |
| `topic:engine` | Generator schemas, mechanics, checks and prompts, plus Profile import and product validation |
| `topic:unitlab` | The local UnitLab client, server, settings and library |
| `topic:security` | Security boundaries, credentials and access protection |

`topic:engine` is broader than the validator. `topic:unitlab` includes server responsibilities as well as the client.
`topic:api` replaces `topic:cli`, because there is no independent operational CLI.
`origin:manga-mayhem` identifies an explicit Manga Mayhem product or source relationship, such as a requested game mechanic.

## Find work

Use `is:issue is:open -label:"status:out-of-scope"` for an ordinary open-issue request. Include excluded proposals only when explicitly requested, including requests for all open issues or parked proposals.

| View | Filter |
| --- | --- |
| Current work | `is:issue is:open label:"status:active"` |
| Triage | `is:issue is:open label:"status:triage"` |
| Owner review | `is:issue is:open label:"review:decision" -label:"status:out-of-scope"` |
| Audit review | `is:issue is:open label:"review:audit" -label:"status:out-of-scope"` |
| Closure review | `is:issue is:open label:"review:closure" -label:"status:out-of-scope"` |
| Later work | `is:issue is:open label:"status:deferred"` |
| Future options | `is:issue is:open label:"status:parked"` |
| Excluded proposals | `is:issue is:open label:"status:out-of-scope"` |
| Paused work | `is:issue is:open label:"status:on-hold"` |

## Changes and migration

This repository owns these instructions and its GitHub label definitions. Keep them complete locally. Another repository's private standard is not required reading.

Preserve recorded decisions when migrating labels. Rename or map a label only when its existing meaning and issue evidence support the replacement. Leave ambiguous assignments unchanged and ask Kyle. A label named `wontfix` alone does not prove owner rejection; inactivity does not prove a pause. Cleanup cannot itself accept, reject, close, pause or resume work. Remove a legacy label only after all uses have been resolved.

Use six hexadecimal digits for colors. Reuse a color when a label has the same name and meaning as an existing label in another repository. Different meanings may share colors. Local additions require an explained recurring need and owner agreement; a color has no acceptance meaning.

Before changing an issue, check its owner decisions, current state, applicable topics, optional type, reasons for any applied review, origin or importance labels, and milestone relevance. Missing optional labels are not a checklist failure.
