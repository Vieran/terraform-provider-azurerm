---
description: "Use when reviewing or authoring Terraform AzureRM provider acceptance tests. Covers test naming, required test set, config style, ImportStep, PreCheck, and breaking-change gating."
applyTo: "internal/**/*_test.go"
---

# AzureRM Provider — Acceptance Test Rules

Reference: [reference-acceptance-testing.md](../../contributing/topics/reference-acceptance-testing.md).

## Package & File Layout

- Test files use the `{resource_or_datasource_file}_test.go` naming and live next to the implementation.
- Package name is the service package **plus `_test` suffix** (e.g. `package aadb2c_test`). Enforced by `make test`. No `Test` suffix on struct names.
- Each resource / data source has its own test struct: `type {Name}Resource struct{}` or `type {Name}DataSource struct{}`. Config helpers are methods on that struct.

## Test Naming

- `TestAcc{ResourceName}_{scenario}` — e.g. `TestAccExampleResource_basic`.
- Group related scenarios with `_{category}_{scenario}` (e.g. `TestAccExampleResource_identity_userAssigned`).

## Required Tests

**Resource** — minimum set:
- `_basic` — only `Required` fields.
- `_requiresImport` — uses `data.RequiresImportErrorStep(r.requiresImport)` and reuses `r.basic(data)`.
- `_complete` — all compatible `Required` + `Optional` fields.
- `_update` — provisions `basic`, updates to `complete`, both wrapped in `ImportStep`s. `ForceNew` fields are not tested this way.

**Data Source** — minimum:
- `_basic` — typically reuses the associated resource's `complete` config and asserts on the `Computed` fields (not the ones the user specified).

More complex resources warrant additional focused tests (enable/disable a block, per-setting round-trip, etc.).

## Extending Existing Schemas

- Add a new Optional resource field to the `complete` test when it is compatible with the other fields.
- Add a new Required resource field to every existing configuration for that resource.
- Use a focused test when a new field or block needs additional setup or prerequisites.
- Adding a field with a default, or changing an existing default, requires manual comparison with the current Terraform state and Azure API behaviour; acceptance tests alone may not detect the breaking change.
- Add each new data-source property to the basic test with an explicit `.Exists()` or `.HasValue(...)` assertion.

## Test Body Patterns

- Use `acceptance.BuildTestData(t, "azurerm_...", "test")` and `data.ResourceTest` / `data.DataSourceTest`.
- Add `data.ImportStep(...)` after **every** step that mutates a resource. Pass field names to ignore (`data.ImportStep("password", "database_primary_key")`) for values Azure doesn't return.
- Use `check.That(data.ResourceName).ExistsInAzure(r)` in resource tests; use `.Key(...).HasValue(...)` for data-source assertions on computed fields.
- Never assert on user-specified fields as an existence check in a data source — a missing resource fails to find it anyway.

## Config Style

- Every config includes `provider "azurerm" { features {} }`.
- **The resource under test is placed last** in the HCL, especially when there are multiple resources.
- Pick the **lowest / cheapest SKU** that exercises the code path.
- Use `data.RandomInteger`, `data.RandomString`, `data.Locations.Primary` / `.Secondary` for uniqueness and cross-region tests.
- Prefer indexed format specifiers (`%[1]d`, `%[2]s`) when a value is reused.
- When a config helper is threaded into `fmt.Sprintf` **only once**, pass it directly instead of assigning a temporary variable.
- `requiresImport` reuses `r.basic(data)` and re-declares the resource as `"import"` referencing `azurerm_..._resource.example.*`.

## PreCheck Helpers

- Global `acceptance.PreCheck(t)` covers auth + location env vars — do not duplicate.
- Test-specific prereqs (e.g. `ARM_TEST_...` env vars) go in a `preCheck(t *testing.T)` method on the test struct, called at the top of each `TestAcc...` that needs it.
- On missing optional prereqs call `t.Skip(...)` / `t.Skipf(...)`, **not** `t.Fatal`, so unrelated tests still pass.
- Place `preCheck` near the tests that call it (typically after the `TestAcc...` functions).

## Breaking-Change Gating

- When a resource is deprecated but still functional in the API, gate the test with:
  ```go
  if features.FivePointOh() {
      t.Skipf("Skipping since `azurerm_...` is deprecated and will be removed in 5.0")
  }
  ```
- When the API no longer works, **delete the test file** instead of skipping.
- When renaming a property, keep one `complete` test using the old name (branched on `features.FivePointOh()` to switch to the new name in major-release mode). Update only the config, not the test case shape, wherever possible.

## Resource Identity & List Resource Tests

- Identity tests are generated via `go:generate go run ../../tools/generator-tests resourceidentity -resource-name {name} -properties "..." [-compare-values "..."] [-no-subscription-id]`. Place the directive under the imports.
  - `-properties` maps 1:1 ID→schema fields (`name,resource_group_name`), remapping with `{id_field}:{schema_field}` when they differ.
  - `-compare-values` handles fields exposed only via a parent resource ID.
  - `-no-subscription-id` for IDs (e.g. management groups) without that segment.
- List resource tests provision N (usually via `count`) instances, then run `Query: true` steps with `querycheck.ExpectLength(...)` / `ExpectLengthAtLeast(...)` against the `list.` address. Skip below `tfversion.Version1_14_0` and use `framework.ProtoV5ProviderFactoriesInit(context.Background(), "azurerm")`.

## Write-Only Attributes

- Tests for WO attributes must gate on Terraform ≥ 1.11 via `tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0")))` and set `ProtoV5ProviderFactories` to `framework.ProtoV5ProviderFactoriesInit(context.Background(), "azurerm")`.
- Cover: create with WO, update WO (bump `_wo_version`), migrate from sensitive → WO, migrate WO → sensitive.
- Use `acceptance.WriteOnlyKeyVaultSecretTemplate(data, secret)` to avoid duplicating prereq HCL.
- Include the trigger field in `ImportStep(...)` ignores (e.g. `data.ImportStep("password_wo_version")`).

## Prohibited

- Reordering existing test cases in an unrelated PR.
- Assertions on the fields the config sets when a simpler `ExistsInAzure` check suffices.
- `t.Fatalf` for optional prerequisites.
- Provisioning resources at the top of the file with the resource-under-test buried in the middle.
