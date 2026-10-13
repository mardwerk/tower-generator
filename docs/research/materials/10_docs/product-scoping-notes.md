# Generator research: product scoping

These are dated source observations and proposals. Preserve their recorded review and rejection status. Current local contracts govern acceptance; these notes neither select implementation nor resume paused development.

## Source observations

Tower Generator's selected pure Go `serve` backend uses REST, SSE and a bounded in-process research queue.
Website and script clients call its API without a browser session; Towerright could call it later.
Its backend owns the portable Markdown Wiki and running research.
Detailed interruption and retention behavior remains open; durable resumption is not a major requirement.


Tower Generator's [product outline](../../../PRODUCT.md#established-direction) owns its selected delivery order and separate Agent Skill and generator modes.
Its old application, Lab, Default Profile and Usopp experiment have been removed, preserving them in Git history.


Tower Generator retains Agent Skill and proper generator modes under its [product scope](../../../PRODUCT.md#established-direction); neither mode requires Manga Mayhem development.

## Product map

| Product | Owns | Typical interfaces |
| --- | --- | --- |
| Tower Generator | Character research, portable Markdown Wiki, Tower authoring and later editor | Selected pure Go `serve` backend with REST, SSE and a bounded in-process research queue; website and script clients, possibly Towerright later |

## Ownership boundaries

Tower Generator owns its research and Wiki workflow, Tower authoring and the later [editor](../../../PRODUCT.md#selected-authoring-editor-boundary).


A potential generator/editor integration uses explicitly supplied data and requires its own review; no game is a generator prerequisite.
Tower Generator's [product outline](../../../PRODUCT.md) records its selected backend direction and pending specification.
Its old Default Profile and Usopp experiment are retired; a new Profile and Manga Mayhem integration remain deferred.

## Dependency boundaries

- Tower Generator clients call its `serve` API without a browser session. Process startup and configuration remain; there is no independent operational CLI.
- Tower Generator's backend owns its portable Markdown Wiki and running research. Detailed interruption and retention behavior remains open; durable resumption is not a major requirement.


Kyle selected pure Go execution, REST responses, SSE progress and a bounded in-process research queue.
REST and SSE concern client communication; the queue controls work admission.
Self-hosted Temporal is not selected for the current backend.
Queue bounds, cancellation, retries and revision handling remain to be specified.
This direction does not establish automatic restart recovery; detailed interruption and retention behavior remains open.

## Repository shape

Tower Generator's retired generator and Lab are preserved in its [pre-reset tag](https://github.com/mardwerk/towerwerk/tree/generator-before-validator-reset-2026-09-29).
The October reset removed the later Python application, web app, Lab, Default Profile and Usopp experiment in [Tower Generator PR #120](https://github.com/mardwerk/towerwerk/pull/120), preserving them in Git history.
Independent research evidence and helpers remain; the product repository owns its new workflow specification.
Historical designs are reference material, and only explicitly retained decisions apply.

## Current order

Towerwerk is standalone and its development is on hold; Manga Mayhem is not a prerequisite.
Manga Mayhem's scope owns game delivery order; Tower Generator's [product scope](../../../PRODUCT.md#established-direction) owns Tool development.

## Where each kind of document belongs

| Kind of document | Home |
| --- | --- |
| Raw imports and historical handoffs | Planning [`99-cold-store/`](../99-cold-store/) |


| Kind of document | Home |
| --- | --- |
| Generic Profile Validator, reusable schemas, Profile Template and binary/Profile bundles | td-profile, including its Profile contract and release packaging |

## Profile targets

Manga Mayhem owns its game Profile task and the preceding gameplay evidence under its scope.
