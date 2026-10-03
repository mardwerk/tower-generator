# API-only reset and backend decision

This record preserves the 2026-10-03 reset and recovery evaluation.
Current behavior and pending choices belong to [PRODUCT.md](PRODUCT.md).

## Completed reset

Kyle approved [PR #120](https://github.com/mardwerk/tower-generator/pull/120), applied directly to main at `b085ba8`.
The old Python product, complete web/Lab tree, Node tooling, Usopp experiment and broken setup are removed.
The Usopp authoring scope in [#85](https://github.com/mardwerk/tower-generator/issues/85) is retired; it no longer receives new generation reviews.

The [latest prototype](https://github.com/mardwerk/tower-generator/tree/ea114b2139abf46318a59f632a58fc035f64a5a7) preserves the retired application and Lab.
[Profile history](PROFILE.md#historical-default-profile) preserves the complete Default Profile and its notices.
[Independent research](research/README.md) records the retained evidence and subsequent archive cleanup.
Local credentials and collected Wiki data were untouched, and known retired local tooling was removed.
[scripts/README.md](../scripts/README.md) owns the helper execution policy.

## Recovery check and outcome

A two-step fake-provider experiment compared Go checkpoints with the Temporal Go SDK.
Processes were killed and restarted; Temporal's Worker and development Service reused their persisted database, and no paid service was called.
Both runners reused a recorded identity response and completed synthesis without repeating identity.
For an unrecorded response, Go stopped with uncertain spend while Temporal's selected retry policy repeated the call.
Neither framework guarantees recovery of an unrecorded paid response; these outcomes reflect different retry policies.

All four cases passed, without testing power loss, Wiki commit safety, throughput or production hosting.
The experiment code, generated reports and downloaded runtime were discarded after Kyle selected the [current backend](PRODUCT.md#selected-transport-and-execution).

## Remaining workflow review

Use the [workflow acceptance review](PRODUCT.md#workflow-acceptance-review) for remaining decisions and their owning issues.
Retirement and documentation repair completed [#117](https://github.com/mardwerk/tower-generator/issues/117); a replacement Profile was not implemented.
