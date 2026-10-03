# Paid-step recovery evaluation

Kyle selected pure Go with REST, SSE and a bounded in-process queue on 2026-10-03.
This evaluation informed that discussion; it does not add durable research execution or Temporal to the product.
The [product rules](../PRODUCT.md) and [architecture review](../REALIGNMENT.md) own the current requirements and pending workflow choices.

## What was executed

The [standalone experiment](../../scripts/recovery-evaluation/README.md) ran two fixed fake paid steps, identity and synthesis, through a standard-library Go checkpoint runner and the actual Temporal Go SDK.
An independent HTTP provider recorded every invocation before responding, without idempotency or result lookup.
Identity responses contained a fresh nonce; synthesis had to consume the exact result returned by the coordinator.
No real provider or paid call was used.

The driver force-killed the Go executable and restarted it against the same store and operation ID.
For Temporal, it force-killed both the Worker and Service, then restarted them against the same development database and Workflow ID.
The provider stayed alive across those interruptions.
Every interrupted process exited through SIGKILL; all four cases passed.

The final observed run began at `2026-10-03T16:15:26.418239Z` on Linux amd64, kernel 7.2.7, with Python 3.14.7 and Go 1.26.8.
Temporal used Go SDK 1.49.0, official CLI 1.9.1 and development Server 1.32.0.
The downloaded CLI matched its official release checksum.
Runtime files, ledgers, development databases and binaries stayed outside Git.

## Recorded-result recovery

The first interruption followed recorded identity completion and preceded any synthesis attempt.
Go had an immutable completed-response checkpoint; Temporal history contained one `ActivityTaskCompleted` and no second Activity schedule.
After restart, both completed synthesis using the original identity nonce without repeating identity.

| Runner | Identity invocations | Synthesis invocations | Outcome |
| --- | --- | --- | --- |
| Go checkpoints | 1 | 1 | Completed; saved identity reused |
| Temporal | 1 | 1 | Completed; recorded Activity result replayed |

This demonstrates recovery of a response that was actually recorded before interruption.
It does not demonstrate recovery of a response the provider produced but the coordinator never saved.

## Uncertain paid-call outcome

The second interruption followed provider success but preceded the coordinator's completion record.
The Go runner had an attempt reservation without a response checkpoint.
Temporal had no completed identity Activity in history.

| Selected demonstration policy | Identity invocations | Synthesis invocations | Outcome |
| --- | --- | --- | --- |
| Go stops when spend is uncertain | 1 | 0 | `spend_uncertain`; no automatic retry |
| Temporal permits at most two Activity attempts | 2 | 1 | Completed after repeating identity |

These are deliberately different policies, not an inherent billing advantage for either runner.
Go could retry and incur another charge; Temporal could stop instead.
Neither can recover an unrecorded provider response without an enforced provider idempotency key or retrievable result.
The product must specify retries and uncertain spend separately from its execution framework.

## Meaning for the selected backend

The two-step checkpoint experiment needs one Go process and files but adds application recovery code.
The Temporal experiment needs application Worker code plus a Service and durable execution history.
Its SQLite development Service is a fixture, not evidence for a production self-hosted deployment.

Kyle stated that restart recovery is not a major requirement and selected pure Go with REST, SSE and an in-process queue.
Do not turn this optional checkpoint experiment into an initial product requirement.
Specify queue admission and limits, cancellation, clarification, progress retention and interruption outcomes with the research workflow.
Coherent Wiki updates and interrupted file-commit recovery remain separate storage questions.

## Reproduction and limits

Run the [driver](../../scripts/recovery-evaluation/run.py) with Python's standard library and an official Temporal CLI:

```sh
python3 -B scripts/recovery-evaluation/run.py \
  --temporal-cli /path/to/temporal \
  --output /tmp/tower-recovery-report.json
```

The driver builds both binaries in a temporary directory and cleans up its processes and local runtime files.
Its report includes version strings, barriers, Service history, ledger counts, actual recovered results and interruption outcomes.
Go-only mode also passed and explicitly reports Temporal as unvalidated.
Focused Go tests, the Go race check, Temporal workflow tests, module builds, vet, link checks and diff hygiene passed.

This evaluates process death on the same machine and storage with tiny responses and one operation per case.
It does not establish power-loss durability, safe Wiki replacement, deployment compatibility, overnight clarification, throughput or total operating cost.
It implements no product API, SSE stream, queue or generation workflow.
