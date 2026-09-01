---
name: azurerm-code-review
description: "Review terraform-provider-azurerm pull requests and code changes. Use for PR review, committed-change review, local diff review, Go implementation review, acceptance-test review, and provider documentation review."
---

# AzureRM Code Review

## Purpose

Run an evidence-based, audit-only review of changes to `terraform-provider-azurerm`.

This skill is the review workflow. Current contributor documentation under `contributing/` is the authority for published provider rules. Files under `references/` contain review methodology and explicitly adopted rules that are not fully codified in the contributor guide.

## Mandatory preflight

Before classifying any finding:

- Read this file completely.
- Resolve the authoritative review scope and inspect every changed file in that scope.
- Read [review-core.md](./references/review-core.md).
- Read [upstream-routing.md](./references/upstream-routing.md), then load every contributor topic selected by its routing table.
- Load the applicable local delta references based on changed paths:
  - `internal/**/*.go`, excluding test-only scope: [implementation-delta.md](./references/implementation-delta.md)
  - `internal/**/*_test.go`: [testing-delta.md](./references/testing-delta.md)
  - `website/docs/**/*.html.markdown`: [docs-delta.md](./references/docs-delta.md)
- Read [review-passes.md](./references/review-passes.md) before freezing findings.
- When structured intermediate records are supported, use [finding-schema.json](./references/finding-schema.json).

Do not begin a normal review when the authoritative change-set cannot be determined. State the missing input instead.

## Audit-only boundary

- Do not modify files or propose patches unless the user explicitly asks for fixes after the review.
- Do not run acceptance tests or broad test suites during normal review unless explicitly requested.
- A review may inspect tests, CI output, and generated output.
- If a required tool is unavailable, report that fact. Never simulate tool output.

## Workflow

### Gather scope

- Prefer the cloud review platform's current pull request file list, patches, base revision, and head revision.
- Treat an explicit pull request identifier as authoritative only when it agrees with the active review context or the user explicitly overrides that context.
- Use a branch diff only when no pull request context exists or the user requests branch-wide review.
- Include added, modified, deleted, renamed, and relevant generated companion files.
- Treat `vendor/**` as non-actionable generated dependency content. Review the dependency or generation input that introduced the change instead.
- Never silently omit a changed file.

### Classify changed files

Classify each file before reviewing it:

- provider implementation
- acceptance or unit test
- provider reference documentation
- contributor or workflow documentation
- generated artifact
- vendored dependency
- other user-visible or operational content

Use this classification to select contributor topics and local delta references.

### Load current repository guidance

- Follow [upstream-routing.md](./references/upstream-routing.md).
- Read the selected contributor files from the checked-out revision being reviewed, not from model memory.
- Inspect the closest same-service implementation and tests when a provider pattern is needed.
- Do not infer a mandatory policy from an unrelated legacy implementation.

### Run the primary review

Review the complete diff for:

- correctness and observable failure paths
- lifecycle, state, update, import, and deletion behavior
- schema and API mapping
- errors, nil values, unknown values, and trust boundaries
- required companion implementation, test, and documentation artifacts
- acceptance-test behavior and embedded Terraform
- documentation parity and runnable examples
- security exposure and secret handling
- cross-platform or sibling-resource consistency when directly relevant

Create an internal candidate record for each possible Issue. A candidate must identify the changed scope, concrete evidence, impact, confidence, and governing source.

### Run internal review passes

Apply [review-passes.md](./references/review-passes.md) in this order:

- primary reviewer
- architect
- skeptic
- advocate, only when candidate Issues exist

The same agent may perform the passes sequentially when subagents are unavailable. Keep intermediate role output internal and publish one deduplicated final review.

### Freeze and publish

- Finish evidence gathering and resolve every candidate before writing the final review.
- Order Issues by severity and impact.
- Put non-blocking design direction or unresolved uncertainty in Observations.
- Keep one concern in one classification only.
- Ensure the merge-readiness verdict agrees with the final Issue set.

## Finding format

Each Issue must include:

- severity
- concise title
- changed file and narrow line location
- evidence from the diff or current repository
- concrete failure, regression, or violated mandatory rule
- impact
- one deterministic correction
- source, using a contributor topic or local delta rule ID

Use Observations for preferences, uncertain concerns, and design direction that lacks a mandatory source. Do not inflate the review with generic praise.

## Final output

Use this order:

1. Issues, highest severity first
2. Observations
3. Validation status
4. Short change summary and merge-readiness assessment

If there are no Issues, say so directly and identify any remaining validation gap. Do not publish internal candidate records, role debate, or skill-loading narration.