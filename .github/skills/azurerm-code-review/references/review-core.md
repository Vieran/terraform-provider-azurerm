# Review Core

## Source precedence

Use sources in this order:

1. The authoritative current diff and changed files.
2. Current repository contributor guidance and pull request template.
3. Current implementation, schema, tests, generators, and same-service patterns.
4. Applicable local delta rules bundled with this skill.
5. Completed tool and CI output.
6. External material only when repository evidence is insufficient.

Published contributor guidance governs provider behavior. Local delta rules fill documented gaps; they do not override a conflicting current upstream rule.

## Evidence rules

### REVIEW-EVID-001: Do not guess

An Issue requires evidence for both the problem and requested correction. If evidence is incomplete, use an Observation or omit the concern.

### REVIEW-EVID-002: Verify rendering artifacts

Confirm apparent encoding, whitespace, truncation, or line-break problems against the actual file before reporting them.

### REVIEW-EVID-003: Attribute policy claims

Do not call a preference mandatory without a contributor topic, adopted local delta rule, completed CI finding, or unambiguous correctness failure.

### REVIEW-EVID-004: Fresh audit

Treat each invocation as a fresh review. Do not reuse an earlier file list, tool result, or review body as current evidence.

### REVIEW-EVID-005: Current-run facts only

Describe what the current evidence establishes. Do not compare against earlier agent runs unless the user asks for that comparison.

## Classification

### REVIEW-CLASS-001: Issues are actual problems

An Issue is a demonstrated defect, regression, security exposure, missing mandatory companion, policy violation, or correctness risk with a concrete failure path.

### REVIEW-CLASS-002: Observations are non-blocking

Use Observations for design preferences, follow-up ideas, uncertainty, or acceptable alternatives.

### REVIEW-CLASS-003: One concern, one classification

Do not duplicate a concern across Issues, Observations, CI findings, or role passes. Merge evidence into one record.

### REVIEW-CLASS-004: Correction is deterministic

Give one concrete correction. Do not list competing solutions unless the user asks for options.

### REVIEW-CLASS-005: Severity follows impact

- `critical`: credential exposure, destructive corruption, or broadly exploitable behavior
- `high`: likely runtime failure, state corruption, unsafe lifecycle behavior, or major regression
- `medium`: constrained correctness failure, missing required coverage, or user-facing contract violation
- `low`: limited defect or mandatory consistency issue with small impact
- `observation`: non-blocking direction or uncertainty

## Scope rules

- Review every changed file in the authoritative scope.
- Classify added, modified, deleted, and renamed files accurately.
- Do not ask contributors to hand-edit `vendor/**`; find the actionable source.
- For generated files, identify whether the changed source and generated output remain aligned.
- Review user-visible text, commands, examples, and error messages in every file type.
- For installer or script pairs, check cross-platform behavior when both implementations are expected to remain aligned.
- For AI customization changes, check source precedence, determinism, duplicated authority, and routing consistency.

## Provider-specific review checks

- A new resource must include Resource Identity, a list resource, list query tests, and list-resource docs unless the documented maintainer exception path applies.
- Do not report a generic missing-list defect when implementation evidence shows a singleton or get-only resource. Require the exception to be justified instead.
- New ephemeral resources require service registration, docs, and Terraform-version-gated tests.
- New provider-defined functions require docs and focused framework tests.
- Update logic must not return after handling one changed field when that would silently skip concurrent changes.
- Create-time existing-resource checks must honor the provider overwrite feature when the applicable contributor guidance requires it.
- Callback-based create flows for resources with Resource Identity must set both ID and identity at the required point.
- A boolean-like string enum is design feedback, not automatically an Issue.
- `validation.StringIsNotEmpty` alone is not automatically an Issue; require evidence that stronger validation is known and mandatory.

## Test execution

- Normal review inspects tests but does not execute acceptance tests unless requested.
- `ImportStep()` is strong broad state validation but does not prohibit targeted checks for behavior import cannot prove.

## Output rules

- Findings lead the review.
- Cite the narrowest useful changed location.
- Explain behavior and impact, not only the rule name.
- Disclose missing evidence and unavailable validation plainly.
- Freeze and deduplicate findings before output.
- A clean final Issue set requires a merge-ready verdict, subject to disclosed validation gaps.
- A non-empty final Issue set cannot receive a merge-ready verdict.