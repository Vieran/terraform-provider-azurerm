# Upstream Contributor Guidance Routing

## Loading rule

Read contributor guidance from the checked-out revision under `contributing/`. Do not copy published rules into review findings from memory.

Always read:

- `contributing/README.md`
- `.github/pull_request_template.md`

Read `contributing/topics/guide-opening-a-pr.md` when reviewing PR structure, generated companions, scope, title, body, or contributor workflow.

## Route by changed surface

| Changed surface or behavior | Contributor topics to read |
| --- | --- |
| Provider implementation under `internal/**/*.go` | `best-practices.md`, `reference-errors.md`, `reference-naming.md`, `guide-resource-ids.md`, `schema-design-considerations.md` |
| New resource | `guide-new-resource.md`, `guide-resource-identity.md`, `guide-list-resource.md`, `reference-acceptance-testing.md`, `reference-documentation-standards.md` |
| New data source | `guide-new-data-source.md`, `reference-acceptance-testing.md`, `reference-documentation-standards.md` |
| New service package or client wiring | `guide-new-service-package.md`, `high-level-overview.md` |
| Existing resource fields | `guide-new-fields-to-resource.md`, `schema-design-considerations.md` |
| Existing data source fields | `guide-new-fields-to-data-source.md`, `schema-design-considerations.md` |
| Resource versus inline schema decision | `guide-new-resource-vs-inline.md`, `schema-design-considerations.md` |
| Write-only attribute | `guide-new-write-only-attribute.md` |
| API or SDK version update | `guide-api-version.md`, `guide-breaking-changes.md` |
| State migration or schema version | `guide-state-migrations.md`, `guide-breaking-changes.md` |
| Acceptance or query tests | `reference-acceptance-testing.md`, `running-the-tests.md` |
| Resource identity | `guide-resource-identity.md`, `guide-resource-ids.md` |
| List resource | `guide-list-resource.md`, `guide-resource-identity.md` |
| Provider reference docs | `reference-documentation-standards.md` |
| Naming-only change | `reference-naming.md` |
| Error handling | `reference-errors.md` |
| Build, generation, or toolchain change | `building-the-provider.md` |
| Debugging instructions | `debugging-the-provider.md` |

## Content-triggered routing

File paths alone are not enough. Also load topics when the diff introduces their behavior:

- Load API version and breaking-change guidance when an SDK bump changes defaults or semantics.
- Load state migration guidance when schema shape, field names, or persisted values change compatibility.
- Load resource ID guidance when IDs are parsed, compared, imported, flattened, or written to state.
- Load schema design guidance for Required, Optional, Computed, ForceNew, collection shape, validation, or `CustomizeDiff` changes.
- Load list-resource and identity guidance when a new managed resource is introduced even if its companion files are absent from the diff.
- Load documentation standards when implementation changes require docs but no docs file is present.

## Evidence use

- Cite the exact topic and section supporting a mandatory finding.
- If contributor guidance is silent, use current implementation evidence and an applicable local delta rule.
- If a local delta conflicts with current contributor guidance, follow the contributor guidance and record the delta conflict for toolkit maintenance.
- Do not infer a current standard solely from old or unrelated service code.