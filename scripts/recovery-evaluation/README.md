# Paid-step restart evaluation

Run the same two paid-step sequence through a narrow Go checkpoint runner and the official Temporal Go SDK.
The fake provider records every invocation as an independent charge; no paid service is called.
This experiment compares selected recovery policies and does not choose the product architecture.

## Run the comparison

Use Python 3 with its standard library, Go compatible with the two modules, and the official [Temporal CLI](https://docs.temporal.io/cli).
The Temporal module currently requires Go 1.26.8; Go may download its pinned toolchain and module dependencies into the usual caches.
No Python dependencies or project virtual environment are required.

From the repository root:

```sh
python3 -B scripts/recovery-evaluation/run.py \
  --temporal-cli /path/to/temporal \
  --output /tmp/tower-recovery-report.json
```

The driver builds both binaries and keeps the provider ledger, operation stores, Service database and process logs in a temporary directory.
It kills and reaps started processes and removes that directory afterward.
The report goes to stdout; `--output` optionally retains another copy at the specified path.
Keep reports and runtime artifacts outside Git.

To test only the Go runner:

```sh
python3 -B scripts/recovery-evaluation/run.py --skip-temporal
```

That report explicitly marks Temporal `not_validated`.
A successful Go-only result does not validate the Temporal comparison.
Optional uv use is `uv run --no-project scripts/recovery-evaluation/run.py` with the same flags.

## Interruption cases

The sequence resolves an identity and synthesizes from that exact identity response.
The provider returns a fresh nonce on every identity invocation and rejects synthesis without its actual identity input.
Its flushed JSONL ledger is independent of the coordinators and has no lookup or idempotency endpoint.

| Case | Interruption boundary | Go checkpoint policy | Temporal demonstration policy |
| --- | --- | --- | --- |
| `checkpoint` | Identity response recorded, before synthesis is attempted | Restart reuses the checkpoint and completes | Service and Worker restart replay the completed Activity and complete |
| `provider` | Provider returned success, before coordinator records completion | Restart reports `spend_uncertain` and makes no further calls | At most two Activity attempts permit a second charge and completion |

Both cases force-kill actual processes and reuse the same operation identifier and local store.
Temporal additionally force-kills and restarts its Service and Worker, using the same persisted development database.
The parent fake provider remains alive through every restart.

Before checkpoint interruption, the driver verifies the Go response file or Temporal history containing one scheduled and completed identity Activity.
It also verifies that synthesis has not been attempted.
The unknown-spend case requires no identity checkpoint or Activity completion before interruption.
After recovery, synthesis must consume the saved identity result, not a fixed fixture.

The expected invocation counts are identity 1 and synthesis 1 for both checkpoint cases.
The selected Go unknown-spend policy expects identity 1 and synthesis 0.
The selected Temporal maximum-attempts-2 policy expects identity 2 and synthesis 1.
These different policies do not establish algorithm superiority or a product retry rule.

## Report contract

The version 1 JSON report contains these fields:

| Field | Type and meaning |
| --- | --- |
| `schema_version` | Integer, currently `1` |
| `recorded_at` | UTC ISO timestamp |
| `versions` | Python, platform and Go strings; Temporal CLI and SDK strings when evaluated |
| `flags` | Explicit skip choice and per-case deadline |
| `scope` | Boolean limits identifying process-restart coverage and excluded claims |
| `temporal_validation` | `not_validated`, `attempted`, `passed` or `failed` |
| `cases` | Engine, interruption case, policy, observed counts, ledger, killed/restarted process IDs and outcome |
| `passed` | Boolean result for the requested subset; also inspect `temporal_validation` |
| `error` | Optional driver setup error; individual cases may also contain an error |

Each case has `engine`, `case`, `operation_id`, `execution_policy` and `passed`.
Successful cases include provider invocation counts and recovered outcomes; completed cases prove synthesis consumed the identity result.
The driver exits nonzero on a setup failure or a failed requested case.
All commands and waits have finite deadlines; failures and Ctrl+C run child-process cleanup.

## Limits

The persisted Temporal development server is an experiment fixture, not a production deployment recommendation.
This tests process restarts with tiny responses and one operation per store.
It does not test power loss, interrupted Wiki commits, deployments, code migrations, overnight waits, provider idempotency or throughput.
It provides no performance or total ownership cost claim.
The [architecture review](../../docs/REALIGNMENT.md) keeps those guarantees and the product choice separate.

See the [Go runner](go-checkpoint/README.md) and [Temporal runner](temporal/README.md) for their implemented policies and focused tests.
