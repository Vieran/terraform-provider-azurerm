---
name: code-review-query-contribution-guide
description: 'Load the AzureRM contribution guide before reviewing code. Use for every code review, pull request review, diff review, or review of changed files in this repository.'
user-invocable: false
disable-model-invocation: false
---

# Query Contribution Guide

Read the repository's complete contribution guidance before producing any code review findings.

## Procedure

1. Recursively enumerate every Markdown file under `contributing/`. Do not rely only on the links currently present in `contributing/README.md`.
2. Read every discovered file in full, including `contributing/README.md` and every file under `contributing/topics/`, even when a topic does not initially appear relevant to the change.
3. Verify that every file from the enumerated list has been read before starting the review. Do not rely on memory, summaries, or assumptions about any part of the guide.
4. Build a checklist from all applicable rules in the complete guide and evaluate every changed file against it.
5. Cite the relevant contribution guide file when reporting a rule violation.

Do not start the substantive review until all discovered contribution guide files have been read. If the guide cannot be fully enumerated or any discovered file cannot be accessed, report that the review is blocked instead of silently skipping part of the guide.