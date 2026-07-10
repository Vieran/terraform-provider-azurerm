---
description: "Use when reviewing or authoring Terraform AzureRM provider documentation (website markdown for resources, data sources, guides). Covers examples, argument/attribute ordering, descriptions, notes, and formatting."
applyTo: "website/**"
---

# AzureRM Provider — Documentation Rules

Reference: [reference-documentation-standards.md](../../contributing/topics/reference-documentation-standards.md).

## Page Structure

- Resource and data-source pages use frontmatter containing `subcategory`, `layout: "azurerm"`, `page_title`, and a concise `description`.
- Resource titles use `# azurerm_{name}`; data-source titles use `# Data Source: azurerm_{name}`.
- Pages contain a short description, Example Usage, Argument Reference, Attributes Reference, and Timeouts. Resource pages also contain Import.

## Examples

- Every resource / data source page must include a working example that passes `terraform plan` when copy-pasted.
- Use `example` as the instance name: `resource "azurerm_resource_group" "example"`.
- Name values should be simple and prefixed with `example-` where the field's naming rules allow (`name = "example-resource-group"`). Otherwise use the simplest valid value.
- Include only the same fields that the resource's basic acceptance test uses (plus needed dependencies). Do not exhaustively document every optional argument.
- Avoid multiple examples unless a scenario is genuinely difficult; complex scenarios belong under `examples/` in the repo.
- **Do not include a `terraform` or `provider` block** in resource / data source examples.
- Use `hcl` code fences for HCL — never `terraform`.
- Use the most specific code-fence language for other snippets.

## Argument Reference

**Ordering** (applies to both typed and untyped, and inside every subsection):
1. Fields making up the resource ID — last user-specified segment first (e.g. `name`, then `resource_group_name` / `parent_resource_id`).
2. `location`.
3. Required, alphabetical.
4. Optional, alphabetical. `tags` is always last.

**Descriptions:**
- Format: `` * `field_name` - (Required|Optional[, Write-Only]) Description. ``
- Keep them concise; move detail into a note.
- `ForceNew: true` fields end with: `Changing this forces a new resource to be created.`
- Enum validation (`StringInSlice`): `` Possible values are `value1`, `value2`, and `value3`. ``
  - Single value: `` The only possible value is `value1`. ``
  - Range (`IntBetween`, `FloatBetween`): `` Possible values range between `1` and `100`. ``
- Default: `` Defaults to `default1`. ``
- Write-only attributes: `(Optional, Write-Only)` in the parentheses; the trigger version field is just `(Optional)`.

**Block arguments** need two entries:
1. Top-level: `` * `block_argument` - (Optional) A `block_argument` as defined below. `` — use the correct indefinite article (`A` / `An`).
2. A subsection after the top-level list; multiple subsections are ordered alphabetically.
   - Order inside a subsection: Required alphabetical, then Optional alphabetical.

## Attributes Reference

- Order: `id` first, then remaining attributes alphabetical.
- Descriptions must **not** include possible values or defaults.
- Block attributes follow the same two-entry pattern (`A ... as defined below.`); subsection order: `id` (if present) then alphabetical.

## Timeouts & Import

- Document every supported operation in a `## Timeouts` section with its default: normally 30 minutes for resource create/update/delete and 5 minutes for resource or data-source read.
- Resource pages include a `## Import` section with a copy-pasteable `terraform import` command in a `shell` code fence and a representative Resource ID.
- Data sources do not have an Import section.

## Notes

- Only these three forms are accepted, and always with the exact `**Note:**` marker:
  - `-> **Note:**` — informational (tips, extra info, links).
  - `~> **Note:**` — warning (must-know to avoid errors that are recoverable).
  - `!> **Note:**` — caution (irreversible / data-loss risk).
- Do not use `Info:`, `Important:`, `Be Aware:`, `NOTE`, or omit the colon.
- **Do not document breaking changes in resource pages.** Minor-version breaking changes go at the top of the changelog; major-version ones go in the upgrade guide (`website/docs/{version}-upgrade-guide.markdown`).
- Do not add "This property will do X in v5.0" notes for feature-flagged future behavior.

## Deprecations & Removals

- When a resource / data source is soft-deprecated, add a note on the page:
  ```markdown
  ~> **Note:** The `azurerm_example` resource has been deprecated because [reason] and will be removed in v5.0 of the AzureRM Provider.
  ```
- When a property is soft-deprecated (renamed): remove the old property from the docs and add the new one. Do **not** add forward-looking notes about default / behavior changes that only apply post-major-release.
- Removals / breaking schema changes are recorded under the appropriate section of `website/docs/{version}-upgrade-guide.markdown`, resources listed alphabetically.

## List Resources

- Docs live under `website/docs/list-resources/`.
- Page frontmatter uses `layout: "azurerm"`, `page_title: "Azure Resource Manager: azurerm_{name}"`, `subcategory: "{Service}"`, and a concise `description`.
- Include: description, one or more `list "azurerm_..."` example blocks, and an Argument Reference for the list config (typically `resource_group_name` optional and `subscription_id` optional with the default note).

## Style

- Documentation is written in International English.
- Tone is helpful and kind, assuming the reader may be unfamiliar with the resource / service.
- Keep the layout / ordering rules consistent across new pages — reviewers rely on it.
