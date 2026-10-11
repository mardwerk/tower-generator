# Agent instructions

Read [README.md](README.md) before changing this repository.
Mardwerk agents with private Planning access must also read its [AGENTS.md](https://github.com/mardwerk/planning/blob/main/AGENTS.md) and master [GLOSSARY.md](https://github.com/mardwerk/planning/blob/main/GLOSSARY.md).
Public contributors do not need Planning access.

When managing issues, labels or milestones, read [local topic definitions](docs/ISSUE_LABELS.md).
Agents with Planning access must also read its [issue labels standard](https://github.com/mardwerk/planning/blob/main/10_docs/ISSUE_LABELS.md), using the local checkout when available.
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

Use this repository's configuration when an engineering skill requests tracker, triage or domain conventions.
Repository rules take precedence over upstream skill defaults.

### Issue tracker

Track product work in GitHub Issues for mardwerk/tower-generator.
Read [tracker conventions](docs/agents/issue-tracker.md) before publishing or updating tracked work.

### Triage labels

Use the existing label lifecycle and [triage compatibility rules](docs/agents/triage-labels.md).
Read these rules before translating an upstream triage role into labels or issue state.

### Domain docs

Use one product context and the existing document owners.
Read [domain conventions](docs/agents/domain.md) before domain exploration or recording terminology and decisions.
