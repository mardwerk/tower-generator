# Issue tracker

Product issues and specifications live in [mardwerk/tower-generator](https://github.com/mardwerk/tower-generator/issues).
Use the gh CLI from this checkout and preserve recorded owner decisions when updating work.

Read an issue with `gh issue view <number> --comments` and inspect its labels and linked contracts.
List relevant work with `gh issue list` and explicit state or label filters.
Use a temporary body file with `--body-file` for multiline issue bodies and comments.

Follow [local labels](../ISSUE_LABELS.md) for issue classification and lifecycle.
Follow [triage compatibility](triage-labels.md) when an upstream skill names a triage role.
Use native issue dependencies when available and explicit linked prerequisites otherwise.
A label, passed check or merged proposal does not accept an unselected contract.

PRs as a request surface: no.
Pull requests carry reviewable changes; they do not form an additional feature-request queue.

When a skill says to publish to the issue tracker, create or update the owning GitHub issue.
When a skill says to fetch a ticket, read that issue and its recorded decisions.
Product behavior belongs here; shared boundaries and vocabulary proposals belong in Planning with links to product evidence.
