---
description: "Routes GitHub Copilot pull request reviews through the AzureRM review skill and the applicable file-scoped rules."
applyTo: "**"
---

# AzureRM Provider - Code Review Routing

When performing GitHub Copilot code review:

- Load and follow `.github/skills/azurerm-code-review/SKILL.md` for review scope, evidence, classification, cross-file completeness, and final verification.
- Apply every path-specific instruction whose `applyTo` pattern matches a changed file. Load all matches when patterns overlap; for example, `*_test.go` files use both the general Go rules and the acceptance-test rules, with the test-specific rule taking precedence for test structure and behavior.
- Treat the skill as the owner of review methodology and the path-specific instructions as the owners of domain rules. Do not duplicate or weaken domain rules in review comments.
- Use current repository contributor guidance and the applicable path-specific instructions as mandatory sources. Do not present a preference or an unsupported pattern as a required change.

This routing file is review-only. It does not prescribe implementation-agent workflow or a fixed prose template for the review.