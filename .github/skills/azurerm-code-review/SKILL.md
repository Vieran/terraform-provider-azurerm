---
name: azurerm-code-review
description: Review pull requests for the Terraform AzureRM Provider. Use for GitHub Copilot code review of any change set to find evidence-backed defects, enforce matching repository instructions, check cross-file companion artifacts, and filter false positives before commenting.
---

# AzureRM Provider Code Review

Use this skill for pull request review. It owns the review method; matching files under `.github/instructions/` own the detailed Go, acceptance-test, and documentation rules.

## Sources And Precedence

Use evidence in this order:

1. The changed files and pull request diff.
2. `contributing/README.md`, relevant files under `contributing/topics/`, and `.github/pull_request_template.md`.
3. Path-specific instructions matching each changed file.
4. Current implementation, tests, generated code, and sibling resources in this repository.
5. Tool output supplied with the review.
6. External references only when repository evidence cannot establish service semantics.

Do not guess about Azure API behavior, SDK model shape, validation constraints, or contributor policy. If evidence is insufficient, omit the claim or make it a non-blocking observation.

GitHub code review is a static audit unless test, linter, or CI results are present in the review context. Do not claim that a command passed or failed unless its result is available. Recommend the narrowest relevant validation when execution evidence is missing, and treat supplied tool output as evidence rather than assuming it is current or complete.

## Review The Complete Change

- Consider every changed file and the interactions between them. Do not silently skip additions, modifications, or deletions.
- Review user-visible text, commands, examples, generated inputs, manifests, scripts, and customization files as well as Go code.
- Verify apparent formatting, encoding, truncation, or line-break defects against the actual file content before reporting them; rendered diffs and review UI wrapping are not sufficient evidence.
- Files under `vendor/**` are non-actionable generated dependencies. Review the dependency or generation input that introduced the change rather than asking for direct vendor edits.
- Keep findings attributable to the pull request. Do not report unrelated pre-existing code unless the change makes it newly reachable or materially worsens it.
- Use the instructions from the pull request base branch; a pull request cannot make its own new review rules govern that same review.

## Four-Pass Method

Perform these passes before posting findings. They are internal review steps, not separate output sections.

### 1. Primary correctness pass

Check the changed behavior against the applicable instructions and nearby implementation. Trace schema, expand/flatten, CRUD, state, import, tests, and docs far enough to identify concrete regressions.

### 2. Architecture pass

Check structural fit: typed versus untyped model, schema shape and naming, resource decomposition, sibling consistency, registration, generated artifacts, and long-term maintainability. Architectural preferences are observations unless a current mandatory source makes them requirements.

### 3. Skeptic pass

Try to produce a concrete failure scenario for each suspicious change. Check:

- correctness, nil and zero-value handling, and error paths;
- ordering, concurrent dirty fields, and partial-update behavior;
- validation and trust boundaries;
- Azure PATCH residual state and omission semantics;
- sensitive data exposure;
- behavior-changing branches without targeted tests;
- drift between related Linux/Windows or otherwise paired implementations.

Do not duplicate an existing concern; strengthen it with better evidence.

### 4. False-positive pass

Challenge every candidate issue from the author's perspective. Search for an existing guarantee in schema validation, callers, SDK behavior, comments, tests, and sibling patterns. Keep the issue only when no evidence-backed defense resolves it. Lower severity when a defense reduces impact, and remove findings that are proven false positives.

## Cross-File Completeness

When relevant to the change, verify the complete companion set:

- New managed resource: implementation, registration, Resource Identity, list resource, resource acceptance tests, list-query tests, resource docs, and list-resource docs.
- List support added to an existing resource: Resource Identity, list implementation, registration, query tests, and docs in the same change.
- List omission: accept only an evidence-backed singleton/get-only case or the documented maintainer exception path; a generated SDK list method alone does not prove meaningful list semantics.
- New ephemeral resource: `sdk.EphemeralResource` implementation, `EphemeralResources()` registration, Terraform 1.10-gated tests, and docs under `website/docs/ephemeral-resources/`.
- New provider-defined function: framework function implementation and registration, Terraform 1.8-gated unit tests, and docs under `website/docs/functions/`.
- Behavior or schema changes: corresponding focused tests, reference docs, changelog, and next-major upgrade guidance where repository policy requires them.

## Other Changed Surfaces

- Bash, PowerShell, installer, or packaging changes: check cross-platform parity, shared manifest use, command help, and shipped asset lists.
- `.github/prompts/**`, `.github/instructions/**`, and `.github/skills/**`: check discovery metadata, path matching, overlapping `applyTo` composition and precedence, duplicate normative rules, stale references, contradictory instructions, and recursion or routing loops.
- User-visible prose: check spelling, grammar, terminology, command plausibility, and a helpful professional tone.

## Finding Bar

An issue must identify all of the following:

- the changed file and relevant line or smallest useful range;
- the concrete defect, regression, missing requirement, or policy violation;
- evidence from the diff, repository, instruction, or tool output;
- the failure scenario or user impact;
- one direct correction path.

Use observations only for non-blocking uncertainty or design direction. Do not turn style preferences into issues. Do not report the same concern in multiple classifications, and do not pad the review with generic praise.

Prioritize findings by impact:

- Critical: security compromise, destructive data loss, or broadly unusable provider behavior.
- High: common-path correctness failure, state corruption, panic, or breaking public behavior.
- Medium: conditional behavioral defect, missing required companion, or meaningful validation/test gap.
- Low: narrow correctness or maintainability problem backed by a repository rule.

If there are no evidence-backed issues, say so plainly and mention only material residual risk or validation gaps. The final merge-readiness statement must agree with the final issue set.