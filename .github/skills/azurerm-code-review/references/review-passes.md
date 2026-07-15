# Internal Review Passes

Run these passes after the primary reviewer has inspected every changed file. They are internal perspectives, not separate user-visible review sections.

## Shared record

Represent every possible finding using `finding-schema.json` when structured state is available. At minimum preserve:

- stable ID
- roles that evaluated it
- title and changed scope
- severity
- evidence
- concrete reasoning or failure path
- confidence
- status
- governing rule references when applicable

Before advocate adjudication, statuses are `candidate` or `observation`. Do not exchange unlabeled prose between passes when structured records are available.

## Architect pass

Purpose: evaluate structural fit and long-term direction, not repeat line-level correctness review.

Inspect:

- schema shape and public naming
- grouping and singular versus plural cardinality
- resource decomposition and singleton modeling
- typed, untyped, and framework surface selection
- cross-resource and sibling-resource consistency
- required identity, list, ephemeral, test, and docs companions
- maintainability and diff readability

Default to Observation. Create a candidate Issue only when a current contributor topic, adopted local delta rule, or concrete correctness failure makes the direction mandatory. The architect never finalizes a candidate.

## Skeptic pass

Purpose: attack the change for defects missed by the primary reviewer.

Inspect:

- correctness and control flow
- nil, zero, unknown, and malformed values
- error handling and partial failures
- concurrency, ordering, retries, polling, and deadlines
- input validation and trust boundaries
- create, read, update, delete, import, and residual state
- security exposure and sensitive data
- untested behavior-changing branches

For every candidate, state `this breaks when ...` and connect that scenario to changed lines. Strengthen an existing record instead of creating a duplicate. If no concrete failure or mandatory source can be shown, use an Observation.

## Advocate pass

Run only when one or more candidate Issues exist. Purpose: defend intentional design and remove false positives before publication.

For each candidate:

- search the PR description, current code, comments, tests, contributor guidance, local delta, and nearby patterns for design intent
- identify existing validation or trust-boundary guarantees
- distinguish an acceptable trade-off from a defect
- reassess severity using demonstrated impact

Resolve every candidate to exactly one outcome:

- `confirmed`: publish as an Issue at justified severity
- `downgraded`: publish as an Issue at lower severity
- `dismissed`: do not publish as an Issue; preserve the rationale internally and publish an Observation only when useful to the author

No candidate may disappear without an outcome. Inconclusive evidence receives the lower justified classification; intent must not be invented.

## Final synthesis

The orchestrator performs final synthesis after the advocate:

- merge duplicates and combine the strongest evidence
- remove internal role notes from user-visible output
- ensure one concern has one classification
- ensure every Issue has a narrow location, impact, deterministic correction, and source
- sort by severity and user impact
- align the final verdict with the final Issue set

No separate moderator pass is required by this skill.