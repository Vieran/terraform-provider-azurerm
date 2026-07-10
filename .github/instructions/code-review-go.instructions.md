---
description: "Use when reviewing or authoring Terraform AzureRM provider Go code (resources, data sources, schema, CRUD, IDs, error handling). Covers schema conventions, resource lifecycle, breaking changes, and code style."
applyTo: "internal/**/*.go"
---

# AzureRM Provider — Code Review Rules

Reference: [contributing/topics/](../../contributing/topics/).

## Resource / Data Source Structure

- **Prefer Typed SDK** (`internal/sdk`) for all new resources and data sources. Untyped (raw `terraform-plugin-sdk`) is legacy.
- Typed resources implement `sdk.Resource` / `sdk.DataSource`; new ones with updates must satisfy `sdk.ResourceWithUpdate`.
- Typed resources and data sources implement `IDValidationFunc()` using the appropriate typed Resource ID validator so imports reject IDs for the wrong resource type.
- File names: `{name}_resource.go`, `{name}_data_source.go`. Registration goes in the package's `registration.go`.
- **Create and Update must be separate methods.** Update must use `d.HasChange(s)` / partial payloads to honor `ignore_changes`.
  - Patch/delta updates: set only changed fields via `pointer.To(...)`.
  - Full updates (`CreateOrUpdate`): `Get` first, nil-check `Model` / `Properties`, then mutate.
- Use `{Operation}ThenPoll` for LROs. Prefer PUT (`CreateOrUpdateThenPoll`) over PATCH (`UpdateThenPoll`) when the caller must clear fields — SDK structs use `omitempty` and cannot send explicit `null` on PATCH.
- Do not introduce `pluginsdk.StateChangeConf`; it is deprecated for service LRO handling. Use an SDK poller or a custom poller.
- In `Create`, check for existing resource with `client.Get` + `response.WasNotFound`; return `metadata.ResourceRequiresImport` when found (unless `SkipImportCheckOnCreateAndAllowOverwritingExistingResources` is set).
- In `Read`, on 404 call `metadata.MarkAsGone(id)` (resources) or return `fmt.Errorf("%s was not found", id)` (data sources).
- Identifier fields (`name`, `resource_group_name`, …) must be sourced from the parsed ID, not the API response.
- Use operation-specific timeouts. The standard defaults are 30 minutes for Create/Update/Delete and 5 minutes for Read, including data-source Read; use a different value only when the service requires it.

## API Versions

- Use a stable ARM API/SDK version. Preview API versions are blocked by `internal/tools/preview-api-version-linter` and are distinct from whether an individual feature is in preview.
- A preview API exception is a last resort: document the compelling reason, service commitment against breaking changes, target stable release date, and responsible contact in `internal/tools/preview-api-version-linter/exceptions.yml`; keep entries sorted by module, service, and version.

## Resource Identity & List Resource (mandatory for new resources)

- New resources must implement `sdk.ResourceWithIdentity` and call `pluginsdk.SetResourceIdentityData` in both `Create` (after `SetID`) and `Read`.
- New resources must ship a companion `{name}_resource_list.go` implementing `sdk.FrameworkListWrappedResource` and be registered via `ListResources()`.
- If Resource Identity is genuinely unsupported (for example, the generator does not support a composite or custom ID), explain why in the PR description.
- If listing is genuinely unsupported, explain why so a maintainer can apply `allow-without-list` / `list-not-supported`.
- Callback variants: use `SetIDAndIdentityCallback` instead of `SetIDCallback` when identity is enabled.
- In a List Resource, if `stream.Result` performs additional API calls after the supplied context is cancelled, capture its deadline before assigning the iterator and create a new context with that deadline inside the iterator.

## Resource IDs

- Use parsers/validators from `hashicorp/go-azure-sdk` first, then `hashicorp/go-azure-helpers/resourcemanager/commonids`. Avoid the legacy `parse/`/`validate/` generators unless already established.
- Composite IDs use `commonids.NewCompositeResourceID` / `ParseCompositeResourceID`.
- **Always parse scoped IDs and IDs returned in API response properties through their typed parser before setting into state** to normalize casing and prevent phantom diffs.

## Schema Rules

**Field ordering** (Arguments + Attributes):
1. ID segments, last user-specified segment first (e.g. `name`, then `resource_group_name` or `parent_resource_id`).
2. `location`.
3. Required, alphabetically.
4. Optional, alphabetically. `tags` is always last.
5. Computed (in typed resources these go in `Attributes()`).

**Naming** (see `reference-naming.md`):
- Match Azure Portal / official marketing terminology when it diverges from the REST API.
- No abbreviations: `virtual_machine` not `vm`, `resource_group_name` not `rg_name`.
- Suffix units: `duration_in_seconds`, `size_in_gb`, `_percentage`, `timestamp_in_utc`. Do not suffix ISO8601 duration fields.
- Certificates/artifacts requiring a format use that suffix, e.g. `certificate_base64`.
- Booleans: append `_enabled` (`compression_enabled`). Flip negatives (`disableFoo` → `foo_enabled`, `no_storage_enabled` → `storage_enabled`). Avoid leading `is_`. Some booleans read better without `_enabled` (`mtls_required`, `terms_of_service_accepted`).
- Blocks: singular names (`rule` block). Lists of primitives and computed-only collections: plural (`allowed_ip_ranges`, `rules`).
- Resource ID references: `{resource_name_without_azurerm_prefix}_id` (e.g. `api_management_id`).

**Validation** (required, not optional):
- All string arguments must have a `ValidateFunc` — at minimum `validation.StringIsNotEmpty`.
- `name` fields validate length and allowed characters (regex + descriptive message).
- Resource IDs use the corresponding `commonids.Validate*ID` / SDK `Validate*ID` function.
- Numeric arguments specify a range (`IntBetween`, `FloatBetween`).
- Validate common formats (IP, port, URL, email, ISO8601).
- Fields with fixed enums use `validation.StringInSlice` with SDK constants (`string(pkg.ConstantValue)`).

**Schema design** (see `schema-design-considerations.md`):
- Portal-required fields → `Required` in Terraform unless the API tolerates omission.
- Arrays: set `MinItems` / `MaxItems` from API constraints.
- `TypeList` block with no required nested fields → optional fields must use `AtLeastOneOf` / `ExactlyOneOf` to prevent empty blocks.
- **Flatten `MaxItems: 1` blocks that contain a single nested property** unless the service team confirms more are imminent (add an inline comment if kept).
- **`Enabled`-only sub-object** → flatten to a single top-level `{feature}_enabled` bool.
- Sub-object with `Enabled` + other required fields → keep as a block with `enabled` inside.
- **Do not expose `None` / `Off` / `Default` sentinel values** for user-configurable fields; omit them from `StringInSlice`, then normalize between empty string and the sentinel in Create/Read. Computed-only fields return the raw value.
- **Discriminator (`type`/`kind`) resources** should be split into per-type resources rather than a single generic one.
- SKU: single arg → top-level `sku` string; multi-arg → `sku` block. Avoid `sku_name` + `sku_family` + `capacity` sprawl for new resources.
- **Preview fields are not supported until GA.**
- Eliminate ambiguity in collection-typed arguments (e.g. no duplicate keys via array-of-objects).
- Prefer Portal terminology and consider grouping related settings into a block for large resources.

**`Optional` + `Computed`** — avoid unless there is no alternative.
- Prefer `Default` (if Azure consistently sets the same value) or making the field `Required`.
- When required, order the flags `Optional`, comment, `Computed`; comment starts with `// NOTE: O+C ` followed by the reason.

**`CustomizeDiff`** using known-after-apply values must use `metadata.ResourceDiff.GetRawConfig().AsValueMap()` and check `.IsNull()` — do not use `Get`/`DecodeDiff` for that comparison. Abstain if the logic requires the unknown value itself.

- Do not return an error for a valid configuration when only a particular state transition cannot be updated in place. Call `metadata.ResourceDiff.ForceNew("field")` from `CustomizeDiff` so Terraform recreates the resource.

## Error Handling

- Wrap errors with context: `fmt.Errorf("verb %s: %+v", id, err)` (`creating`, `updating`, `retrieving`, `deleting`, `waiting for %s to finish provisioning`). Use the parsed ID type — not the raw string — as the format arg.
- Do not start wrapped messages with `failed`, `error`, or a capital letter; a caller adds the prefix.
- Wrap in every function that has two or more error checks.
- Internal errors (impossible states) prefix with `internal-error`: `errors.New("internal-error: context had no deadline")`.
- Use `errors.New` for static messages; `fmt.Errorf` only when formatting.
- When re-returning parse errors (`ParseXxxID`), return `err` directly.
- Wrap argument names and internal schema terms such as `model` and `properties` in backticks in error messages.

## Pointer Helpers

- `pointer.From(x)` instead of manual `if x != nil { … *x }`.
- `pointer.To(v)` instead of temporary vars for addresses.
- `pointer.ToEnum[EnumType](s)` for enum type conversions.

## File Header

New Go files begin (no preceding blank line) with exactly:
```go
// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0
```

## Breaking Changes & Deprecation

- Never break surface in a minor release. Gate breaking changes behind the next-major feature flag (currently `features.FivePointOh()`).
- Update the whole schema block inside the flag, not inline mutations. No anonymous functions on schema fields.
- Renaming a property: keep both, both `Optional + Computed`, mutually `ConflictsWith`, old one gets `Deprecated:` message referencing the new one. In typed models, tag the old field `,removedInNextMajorVersion`; new-only fields tagged `,addedInNextMajorVersion`.
- Removing a resource: add `DeprecationMessage()` (or `DeprecatedInFavourOfResource()`) and conditionally register behind the feature flag.
- Changing a default: gate the new default behind the feature flag with the old default in the `!features.FivePointOh()` branch.
- When adding a property that Azure fills in server-side, set a `Default` matching the API so `terraform plan` is clean.
- **Do not reorder existing properties in the same PR.**

## Write-Only Attributes

- Write-only attributes require Terraform 1.11 or later.
- Pair `foo_wo` (`WriteOnly: true`) with `foo_wo_version` int trigger. Use `RequiredWith`/`ConflictsWith` to keep `foo`, `foo_wo`, `foo_wo_version` consistent.
- Read via `pluginsdk.GetWriteOnly(metadata.ResourceData, "foo_wo", cty.String)`; write in Create and when `foo_wo_version` changes in Update.
- Persist only `foo_wo_version` into state in Read.
- Original sensitive field must not be `ForceNew`, `Computed`, inside a nested collection, a block/map, or on a data source/provider.
- Do not currently add `pluginSdkValidation.PreferWriteOnlyAttribute`; users cannot dismiss its warning without migrating to the write-only attribute.

## State Migrations

- Live under `internal/services/{svc}/migration/` as `{resource}_v{n}_to_v{n+1}.go`.
- Schema in the migration must strip `Default`, `ValidateFunc`, `ForceNew`, `MinItems`/`MaxItems`, `AtLeastOneOf`, `ConflictsWith`, `ExactlyOneOf`, `RequiredWith`, and inline any function-returned elements or feature-flag branches.
- Hook up via `sdk.StateUpgradeData` (typed) or `SchemaVersion` + `StateUpgraders` (untyped). One-way — thorough manual testing required.

## Registration

- Register new resources/data sources/list resources in the service's `registration.go`.
- If a package's `Registration` doesn't yet satisfy `sdk.TypedServiceRegistration` / `FrameworkServiceRegistration`, add the interface and wire it into `internal/provider/services.go`.
