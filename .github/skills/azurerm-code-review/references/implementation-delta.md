# Implementation Review Delta

These rules supplement current contributor guidance. `Inferred maintainer convention` rules require supporting current code or review evidence. `Local safeguard` rules are explicitly adopted by this skill to prevent known review and CI failures.

## Workflow and companions

### IMPL-WF-002A: List retrofits include the full companion set

- Provenance: Inferred maintainer convention.
- When adding list support to an existing resource, require Resource Identity, `*_resource_list.go`, service registration, query tests, and list-resource docs in the same workflow.
- Missing companions are an Issue unless the diff or PR uses the documented maintainer exception path.

### IMPL-WF-004: Ephemeral resources use the framework surface

- Provenance: Inferred maintainer convention.
- Expect `*_ephemeral.go`, `sdk.EphemeralResource`, `Open`, `EphemeralResources()` registration, `website/docs/ephemeral-resources/`, and focused version-gated tests.

### IMPL-WF-005: Provider-defined functions use the provider function surface

- Provenance: Inferred maintainer convention.
- Expect implementation under `internal/provider/function/`, `Metadata`, `Definition`, `Run`, docs under `website/docs/functions/`, and focused framework unit tests.

## Schema and state

### IMPL-SCHEMA-005: Keep bespoke validators service-local

- Provenance: Local safeguard.
- Reuse shared validators when they express the real constraint.
- A new or materially changed bespoke validator should live under the owning service's `validate/` directory with a matching unit test.
- A short one-off composition may remain inline. Do not demand churn of untouched legacy validators.

### IMPL-SCHEMA-013: Guard unknown raw configuration values

- Provenance: Local safeguard extending the upstream `GetRawConfig()` pattern.
- Before calling `LengthInt()`, `AsValueSlice()`, `AsValueMap()`, `Index()`, or equivalent shape inspection on `cty.Value`, check `IsKnown()`.
- Diff-time validation must defer when a required value is unknown rather than panic or treat unknown as empty.
- Do not use Azure SDK enum-pointer helpers on Terraform values inside `CustomizeDiff`.

### IMPL-UPDATE-001: Preserve concurrent field changes

- Provenance: Local safeguard.
- An update branch that handles one changed field and returns early is an Issue when another updatable field can change in the same plan and remain unapplied.
- Accept an early return only when the operation is proven to include every dirty updatable field.

## Errors and logging

### IMPL-ERR-002: Avoid redundant wrapping of comprehensive parser errors

- Provenance: Inferred maintainer convention.
- Return a resource ID parser error directly when it already gives complete user-facing context.
- Add wrapping only when it contributes materially new information.

### IMPL-CODE-001: Comments explain irreducible context

- Provenance: Local safeguard.
- Flag comments that merely narrate assignments, struct initialization, routine nil checks, standard CRUD flow, or obvious field mappings when newly added.
- Keep comments for non-obvious Azure behavior, SDK workarounds, complex logic that cannot be simplified, and non-obvious state management.

### IMPL-CODE-002: Avoid redundant lifecycle logging

- Provenance: Inferred maintainer convention.
- Generic `Creating`, `Reading`, `Updating`, `Deleting`, or `Import check` logging is an Issue when it duplicates Terraform or provider-native lifecycle logs.
- Allow targeted diagnostics that add distinct value, including established not-found and removing-from-state messages.

## Feature-gated behavior

### IMPL-TEST-003: Cover non-default lifecycle branches

- Provenance: Local safeguard.
- When a provider feature changes create, update, delete, import, overwrite, or destroy semantics, require the smallest feasible test for the changed non-default branch.
- Do not accept default lifecycle tests as proof of an unexecuted feature branch.
- Use existing client callback setup patterns for pre-existing remote objects when applicable.

## Review restraint

- Do not make boolean-toggle preferences blocking without a mandatory source.
- Do not require stronger validation merely because only `StringIsNotEmpty` is present; prove that the real constraint is known.
- Do not apply migration, security, caching, concurrency, or API fallback examples from generic companion guides unless current provider evidence supports them.