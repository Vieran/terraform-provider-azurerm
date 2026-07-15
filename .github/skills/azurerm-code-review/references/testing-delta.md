# Testing Review Delta

Use these rules after reading the current acceptance-testing contributor guidance and nearby same-resource tests.

## Coverage selection

### TEST-PATTERN-005: Simple property validation belongs in unit tests

- Provenance: Inferred maintainer convention.
- Do not require an acceptance test solely for simple property validation already covered by a focused unit test.
- Require acceptance coverage only when provider planning, lifecycle, API, or cross-field behavior cannot be proven at unit level.

### TEST-PATTERN-006: Exercise `CustomizeDiff` failures

- Provenance: Local safeguard.
- Require a targeted `ExpectError` acceptance scenario for new or materially changed `CustomizeDiff` constraints when no lower-level harness executes the real diff behavior.
- Existing lifecycle scenarios may cover the valid path; do not duplicate them without a distinct assertion.

### TEST-PATTERN-011: Exercise feature-gated lifecycle branches

- Provenance: Local safeguard.
- Require focused coverage when a non-default provider feature changes CRUD, import, overwrite, or destroy behavior and existing harness patterns can exercise it.
- Prefer one representative test for shared behavior instead of duplicating it across equivalent sibling resources.

## Helper and generator consistency

### TEST-PATTERN-007: Inline one-use config helper calls

- Provenance: Inferred maintainer convention.
- In `fmt.Sprintf` config builders, avoid a local variable used only to pass one nested helper result immediately into `fmt.Sprintf`.
- Keep a local when reused, transformed, or materially clearer.

### TEST-PATTERN-008: Keep helper type names canonical

- Provenance: Local safeguard.
- Use one established helper type for a Terraform resource or data source across main, list, and generated identity tests.
- For new identity-enabled surfaces, verify the name produced by the generator.
- Do not hand-edit generated tests, add alias adapters, or create file-specific helper type variants to hide naming drift.

### TEST-PATTERN-009: Data source setup prefers the associated complete config

- Provenance: Inferred maintainer convention.
- Prefer the associated resource's `complete(data)` helper when it exists and the data source asserts the broader computed surface.
- Accept `basic(data)` or a focused setup when complete configuration adds unrelated coupling or the scenario is intentionally narrow.

## Embedded Terraform

### TEST-PATTERN-010: Embedded Terraform uses repository formatting

- Provenance: Local safeguard.
- Terraform inside Go raw strings must use two-space indentation and no tab indentation.
- Inspect the raw Terraform directly; Go formatting and Go tests do not prove this requirement.

## Azure setup callbacks

### TEST-PATTERN-012: Pollers require deadline-bearing contexts

- Provenance: Local safeguard.
- In `CheckWithClientForResource`, `CheckWithClientWithoutResource`, or `CheckWithClient` callbacks, wrap the supplied context with a timeout or deadline before calling Azure `*ThenPoll` helpers.
- Directly passing a context without a deadline can fail with `the context used must have a deadline attached for polling purposes`.
- Keep quota-sensitive failures separate; they may require sequential execution rather than a timeout change.

## Specialized framework tests

### TEST-WF-004: Ephemeral resource tests

- Provenance: Inferred maintainer convention.
- Expect `*_ephemeral_test.go`, the framework provider factories, the current Terraform minimum-version gate, and an established result assertion pattern.

### TEST-WF-005: Provider-defined function tests

- Provenance: Inferred maintainer convention.
- Expect focused tests under `internal/provider/function/*_test.go`, framework provider factories, the current Terraform minimum-version gate, and assertions against `provider::azurerm::<name>(...)` output.