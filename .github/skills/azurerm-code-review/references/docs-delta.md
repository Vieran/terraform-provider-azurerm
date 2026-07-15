# Documentation Review Delta

Apply these rules only to provider reference pages under `website/docs/**/*.html.markdown`. Read `contributing/topics/reference-documentation-standards.md` first.

Local safeguard rules are explicitly adopted by this skill. Inferred maintainer conventions require current repository or review evidence.

## Block structure

### DOCS-SHAPE-007: Directional block references match position

- Provenance: Local safeguard.
- Use `as defined above` when the referenced subsection is earlier in the same section and `as defined below` when it is later.
- Top-level block bullets continue to use the upstream canonical `as defined below` pattern.

### DOCS-SHAPE-008: Separate block subsections

- Provenance: Local safeguard.
- Place `---` before the first nested block subsection after top-level bullets and between adjacent block subsections.
- Do not place separators between ordinary argument or attribute bullets.

## Runnable examples

### DOCS-EX-003: Resource examples are self-contained

- Provenance: Local safeguard.
- Every Terraform resource, data source, and module reference used by a resource example must be declared on the same page.
- Fix missing references by declaring the dependency, not by deleting the example.

### DOCS-EX-020: Self-containedness is transitive

- Provenance: Local safeguard.
- Added dependencies must include their required inputs, and every new reference they introduce must also resolve on the page.
- Continue until the example is runnable. If evidence is insufficient, record an Observation instead of inventing configuration.

### DOCS-EX-019: Do not invent literals to replace references

- Provenance: Local safeguard.
- Do not replace a Terraform reference with a plausible literal merely to make an example look self-contained.
- A literal replacement is acceptable only when implementation evidence proves the form and the example intentionally teaches a literal scenario.

### DOCS-EX-004: Preserve required `depends_on`

- Provenance: Local safeguard.
- Preserve existing `depends_on` entries and the referenced objects when rewriting an example.
- Add missing declarations instead of shortening the dependency list.

### DOCS-EX-017: Do not invent `depends_on`

- Provenance: Local safeguard.
- Add `depends_on` only when implementation or documentation evidence proves an ordering requirement.

### DOCS-EX-018: Preserve example-adjacent notes

- Provenance: Local safeguard.
- Preserve notes that explain sequencing or validation near an example. Rewrite a note only when the evidence-backed example behavior changed.

### DOCS-EX-012: Keep examples as Terraform

- Provenance: Local safeguard.
- Do not delete an Example section or replace its Terraform block with prose as a repair strategy.

### DOCS-EX-015: Derive example names deterministically

- Provenance: Local safeguard; nit-level only.
- Prefer `example-<full-type-suffix>` for resources and `existing-<full-type-suffix>` for data sources, replacing underscores with hyphens.
- Apply the owning block's full type suffix independently to auxiliary resources.
- Schema validation wins. If a compliant deterministic value cannot be proven, do not guess and do not raise an Issue.

## Concision and notes

### DOCS-NOTE-009: Keep non-resource field documentation concise

- Provenance: Local safeguard.
- Data source, list-resource, ephemeral-resource, and function field entries must explain the field without field-level note blocks.
- Top-level runtime or version notes remain allowed when supported by the page type.

### DOCS-ARG-011: Limit argument bullets

- Provenance: Local safeguard.
- Prefer one sentence and allow at most two concise sentences for an argument definition.
- Keep possible values and defaults in the bullet when readable. Move resource-only caveats into the correctly marked adjacent note.

## Wording

### DOCS-WORD-006: Use the canonical resource name

- Provenance: Local safeguard.
- In Attributes, Timeouts, and Import prose, refer to the exact resource noun phrase established by the page summary rather than a broader service object.

### DOCS-WORD-007: Capitalize the Azure object name `Resource Group`

- Provenance: Inferred maintainer convention.
- Use `Resource Group` when prose refers to the Azure object, including common `resource_group_name` field descriptions.
- Do not capitalize generic prose about grouping resources.

## Deprecation review

### DOCS-DEPR-001: Review the vNext surface

- When schema and feature-gate evidence identify a next-major surface, current reference docs should focus on that vNext surface.
- Do not require legacy-only fields to remain in live reference documentation.

### DOCS-DEPR-002: Exclude legacy-only fields

- A field proven to exist only on the non-vNext path must not appear in current Arguments or Attributes sections.
- Require implementation evidence before classifying a field as legacy-only.

## Review restraint

- Documentation Issues require an exact upstream rule, one of the local rule IDs above, or an unambiguous functional failure.
- Naming derivation and Oxford-comma feedback are low priority and must not make an otherwise valid page unmergeable.
- Data source examples may assume the looked-up object already exists; do not force resource scaffolding into them.