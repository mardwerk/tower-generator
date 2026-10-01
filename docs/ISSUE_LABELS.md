# Issue labels

The [repository label list](https://github.com/mardwerk/tower-generator/labels) shows the installed labels and their descriptions.
Labels describe work and recorded owner decisions. Applying a label does not accept a specification or select work for the owner.
Maintainers with the required GitHub permissions assign labels. Contributors can describe uncertain scope or classification in the issue body.

| Family | Meaning |
| --- | --- |
| `status:` | The recorded work state, or a reason an issue was closed |
| `topic:` | The main responsibility involved |
| `type:` | The kind of change, such as a bug fix, feature, documentation or maintenance |
| `review:` | A pending owner decision or a fresh audit of an existing decision or artifact |
| `origin:` | An explicit relationship to another product or source repository |
| `importance:` | Importance explicitly recorded by the owner |

The topics below describe this repository's responsibilities, including retained future work.

| Topic | Meaning |
| --- | --- |
| `topic:repository` | Repository organization, contributor rules, tracking and governance |
| `topic:operations` | Builds, test environments, publishing, deployment and operational data |
| `topic:interface` | User controls, presentation and interaction |
| `topic:cli` | CLI commands, flags and serve API |
| `topic:engine` | Generator schemas, mechanics, checks and prompts, plus Profile import and product validation |
| `topic:unitlab` | The local UnitLab client, server, settings and library |
| `topic:security` | Security boundaries, credentials and access protection |

`topic:engine` is broader than the validator. `topic:unitlab` includes server responsibilities as well as the client.
`origin:manga-mayhem` identifies an explicit Manga Mayhem product or source relationship, such as a requested game mechanic.
