# Temporal recovery experiment

This standalone module evaluates recovery with the real Temporal Go SDK and development Service.
It is an independent helper, not the Tower Generator backend or an architecture decision.
The parent [driver](../run.py) owns the fake provider, independent spend ledger and process interruptions.
No paid provider is called.

Run the SDK workflow test and build a binary outside the repository:

```sh
go test ./...
go build -buildvcs=false -o /tmp/recovery-temporal .
```

The tested toolchain is Go 1.26.8, Go SDK 1.49.0, Temporal CLI 1.9.1 and its development Server 1.32.0.
The module pins dependency versions in `go.mod` and `go.sum`.
Use an official [Temporal CLI release](https://github.com/temporalio/cli/releases) and verify its release checksum before running the driver.
Keep the downloaded executable, development database, markers and logs outside Git.

Start the development Service with an explicit persistent database and inspect its health:

```sh
temporal server start-dev --ip 127.0.0.1 --port 7233 --headless --db-filename /tmp/recovery-temporal.db
temporal --address 127.0.0.1:7233 operator cluster health
```

Commands are `worker`, `start`, `status`, `history` and `continue`.
Common flags are `--address HOST:PORT` and `--operation ID`; `worker` does not require an operation ID.
`start` also requires `--provider URL` and accepts `--pause-after checkpoint|provider`, `--pause-step identity|synthesis`, and `--marker PATH`.
The `provider` failpoint requires a marker path.
Client commands have an eight-second connection and request deadline.

The Workflow calls an identity Activity, then passes its actual returned result into the synthesis Activity.
The fake provider gives each identity invocation a different result, so the driver can prove that synthesis consumed the retained response.
`status` returns JSON containing `state`, `step`, `checkpoint`, `completed_steps` and `results`.
`history` returns event IDs and types, including Activity scheduling and completion.

The `checkpoint` failpoint waits for a `continue` Signal in Workflow code after `Activity.Get` returns.
Before killing the Worker and Service, the driver requires one recorded `ActivityTaskCompleted` and no second `ActivityTaskScheduled`.
It restarts both processes against the same database and Workflow ID, sends the Signal and verifies one identity call and one synthesis call.

The `provider` failpoint writes a marker inside the first Activity attempt after provider HTTP success, then blocks before reporting completion.
The driver kills both processes, retains the same database and starts replacements.
The explicit demonstration policy allows two attempts, with a 15-second Start-To-Close timeout, a 60-second Schedule-To-Close timeout and a 100-millisecond retry delay.
The second attempt repeats the provider call and completes, exposing a possible duplicate charge under this selected policy.
This does not establish that every Temporal application must retry an uncertain paid call.

The fake provider has no idempotency or result lookup API.
Temporal cannot recover a provider response that was never recorded; enforced provider idempotency, retrievable operations or an explicit uncertain-spend policy require separate evaluation.
This experiment does not test Wiki commits, power loss, human review, workflow upgrades, several Workers, account-wide budgets or production deployment.

The [embedded SQLite Service](https://docs.temporal.io/self-hosted-guide/embedded-server) is for development and testing.
Production Temporal requires an operated Service or Temporal Cloud; the experiment makes no production deployment claim.
See [Activity idempotency](https://docs.temporal.io/activity-definition#idempotency) and [Activity timeouts](https://docs.temporal.io/develop/go/activities/timeouts) for the recovery behavior under test.
