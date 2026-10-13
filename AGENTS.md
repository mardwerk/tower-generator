# Agent instructions

Read [README.md](README.md) and [RULES.md](RULES.md) before changing this repository.

When managing issues, labels or milestones, read [local topic definitions](docs/ISSUE_LABELS.md).
Derive work state from explicit, recorded owner decisions.
Relabelling alone does not select work or accept a specification.

Follow [product rules](docs/PRODUCT.md), [research rules](docs/RESEARCH.md) and the [workflow acceptance review](docs/PRODUCT.md#workflow-acceptance-review).
Use the [realignment record](docs/REALIGNMENT.md) when consulting retired implementation evidence.
Do not present historical commands, schemas or proposals as the new backend contract.

For Profile integration or BTD6 data, follow [Profile ownership and attribution](docs/PROFILE.md).
When generation resumes, review new mappings and generation rules with Kyle in the issue that owns the proposed change before implementation.

Keep downloaded bundles, runtime binaries and generated game-data out of Git.
Put progress and skip reasons in the terminal rather than Tower JSON.

For independent helpers, follow [scripts/README.md](scripts/README.md).
Keep references to the retained [research documents](docs/research/README.md) current.

Wrap prose so each physical line contains at most two sentences.
Put explanatory text between a heading and any subheading.
Before committing, run `git diff --check` and check changed links.
Run checks appropriate to executable changes; document what each check establishes.

## Agent skills

### Issue tracker

Track work in this repository's GitHub Issues. Before tracker operations, read [tracker conventions](docs/agents/issue-tracker.md).

### Triage labels

Keep the existing labels. Before triage, read [role mappings](docs/agents/triage-labels.md).

### Domain docs

Use the local glossary and `docs/adr/` for one context. Before domain exploration or terminology changes, read [domain conventions](docs/agents/domain.md) and [repository rules](RULES.md).
