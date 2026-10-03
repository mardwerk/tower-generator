# Go checkpoint experiment

This standalone experiment runs two fixed fake paid steps, identity and synthesis.
It evaluates restart recovery and does not select the product architecture or paid-call retry policy.
It uses the Go standard library and its own module; it is not the product backend or operational CLI.

Run from this directory with an external temporary store:

```sh
go run . --store /tmp/tower-recovery-store --provider http://127.0.0.1:PORT --operation example
```

The fake provider accepts `POST /paid-step` with `operation_id` and `step` strings.
It returns those strings and a `result` string; the runner never reads the provider's independent ledger.
Synthesis also receives an `input` string containing the actual saved identity result.

The runner reserves each attempt before HTTP and atomically saves its successful response before starting the next step.
It reloads the saved identity result after restart and preserves that checkpoint unchanged.
An attempt without a completed checkpoint produces `spend_uncertain`, exits nonzero and refuses further provider work.
This conservative experiment policy also applies to transport failures and invalid provider responses.

To pause after saving identity, before reserving synthesis:

```sh
go run . --store /tmp/tower-recovery-store --provider http://127.0.0.1:PORT --operation checkpoint \
  --pause-after checkpoint --pause-step identity --marker /tmp/tower-recovery-marker
```

For the uncertain-response case, use `--pause-after provider` instead.
That pause follows HTTP success and response validation but precedes saving the response.
Both pauses write the marker and block until termination.
Build a temporary executable when killing the runner, since killing a `go run` wrapper does not identify its child reliably.
Restart that executable with the same store and operation ID, without the pause flags.

Final stdout is one JSON object with `schema_version`, `operation_id`, `status`, `results` and `reserved_attempts`.
Errors add an `error` string and exit nonzero; completed operations exit zero.
Each operation directory contains versioned `state.json`, immutable attempt reservations and immutable response checkpoints.
On restart, the attempt and checkpoint files reconcile a stale state snapshot.

Run `go test ./...` for checkpoint reuse, a repeated completed operation and refusal of an unknown attempt.
The shared evaluation harness separately kills and restarts the executable against the independent provider ledger.

File writes use a temporary file, `Sync` and rename on the same filesystem.
The experiment tests process termination on the tested platform, not power loss or universal filesystem durability.
It requires one runner per operation store and does not implement locking, queue admission, cancellation, clarification, retention, schema migration or Wiki commit recovery.
Unknown record versions are refused rather than migrated.
Store data and compiled binaries outside Git.
