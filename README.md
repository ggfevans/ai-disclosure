# AI Disclosure Convention

A lightweight, machine-readable convention for declaring AI involvement in source code, designed for individual developers and small open-source projects.

This convention assembles two stable upstream pieces: the [W3C AI Content Disclosure](https://www.w3.org/community/ai-content-disclosure/) vocabulary and the [SPDX](https://spdx.dev/) line-tag format. Together they give you a single pattern that works across any programming language.

- **Version:** 0.1 (2026-05-17)
- **Status:** Draft. Feedback welcome via issues.
- **License:** [CC0 1.0 Universal](./LICENSE). Fork, adopt, or modify without attribution.

## Why this exists

Modern open-source code is increasingly written with AI assistance, ranging from light autocomplete to fully generated modules. The social signals we historically relied on to assess code (commit author, maintainer reputation, line-by-line review) no longer reliably indicate human authorship or oversight. This convention provides a transparent, machine-readable signal so that consumers of the code (humans, linters, CI systems, package managers, license checkers) can apply their own policies.

The convention does not assess whether AI-assisted code is good or bad. It only documents how the code was produced.

This convention also does not satisfy the [EU AI Act Article 50](https://eur-lex.europa.eu/eli/reg/2024/1689/oj), which mandates machine-readable disclosure of AI-generated text content from August 2026. Whether source code falls under that mandate is an open legal question. Voluntary disclosure is sensible regardless of the legal answer.

## The vocabulary

Four values, taken verbatim from the W3C AI Content Disclosure vocabulary:

| Value | Meaning |
| --- | --- |
| `none` | No AI involvement. Use when you want to positively assert human authorship. |
| `ai-assisted` | Human-authored; AI edited, refined, or filled in boilerplate. |
| `ai-generated` | AI-generated with human prompting and review. |
| `autonomous` | AI-generated without meaningful human oversight. |

Absence of a tag means "unknown", not "none". This asymmetry is deliberate. Tagging is a positive act of disclosure, and untagged legacy code is not lying about itself.

## File-level disclosure (SPDX-style)

Add a single-line tag in a top-of-file comment, using the SPDX convention familiar from `SPDX-License-Identifier`. The field names are the same across all languages; only the comment syntax changes.

See [`examples/`](./examples) for full files in JavaScript, TypeScript, Python, and Go.

### JavaScript / TypeScript

```javascript
/**
 * SPDX-AI-Disclosure: ai-assisted
 * SPDX-AI-Model: claude-opus-4-6
 * SPDX-AI-Provider: Anthropic
 * SPDX-AI-Scope: Type definitions AI-generated; logic by hand.
 */

import { foo } from './bar';
```

### Python

```python
# SPDX-AI-Disclosure: ai-generated
# SPDX-AI-Model: claude-opus-4-6
# SPDX-AI-Provider: Anthropic
# SPDX-AI-Scope: Implementation only; tests hand-written.

"""Module docstring."""
```

### Go

```go
// SPDX-AI-Disclosure: ai-generated
// SPDX-AI-Model: claude-opus-4-6
// SPDX-AI-Provider: Anthropic
// SPDX-AI-Scope: Generated boilerplate; reviewed manually.

package mypackage
```

### Any other language

Use the language's top-of-file comment syntax with the same field names:

| Field | Required? | Meaning |
| --- | --- | --- |
| `SPDX-AI-Disclosure:` | yes | One of the four vocabulary values. |
| `SPDX-AI-Model:` | optional | Model identifier (e.g. `claude-opus-4-6`, `gpt-5`). |
| `SPDX-AI-Provider:` | optional | Provider name (e.g. `Anthropic`, `OpenAI`). |
| `SPDX-AI-Scope:` | optional | Free-text description of what the AI did. |
| `SPDX-AI-Date:` | optional | ISO date of the last AI involvement. |

`SPDX-AI-Model:` and `SPDX-AI-Provider:` may be repeated when multiple models contributed.

## Repo-level disclosure

Place an `AI_DISCLOSURE.md` at the repo root with YAML frontmatter. A copy-paste-ready template lives at [`examples/AI_DISCLOSURE.md`](./examples/AI_DISCLOSURE.md):

```yaml
---
disclosure-default: ai-assisted
models-used:
  - claude-opus-4-6
  - gpt-5
providers:
  - Anthropic
  - OpenAI
scope: |
  Most application code is AI-assisted with manual review.
  Tests and migrations may be ai-generated.
  Documentation is human-written unless headers say otherwise.
last-updated: 2026-05-17
---

# AI Disclosure

This repository follows the
[ai-disclosure convention](https://github.com/ggfevans/ai-disclosure).
See per-file headers for overrides.
```

The frontmatter is the machine-readable part; the markdown body is for humans.

## Inheritance and overrides

Disclosure follows a nearest-ancestor model. Specificity wins, in this order:

1. If a file has a `SPDX-AI-Disclosure:` tag, that tag wins.
2. Otherwise, the repo-level `disclosure-default` from `AI_DISCLOSURE.md` applies.
3. Otherwise, disclosure is unknown.

This is the same mental model as CSS specificity, or as frontmatter property inheritance in note-taking tools like Obsidian.

## The "current state" principle

Disclosure reflects the current state of the code, not its history.

If you take AI-generated code and rewrite it substantially during review, downgrade `ai-generated` to `ai-assisted`. If you rewrite it enough to fully claim authorship, remove the header. A human who has thoroughly reviewed, understood, and rewritten a piece of code may reasonably call it their own. This is a moral and pragmatic principle, not a legal one.

Git history is a poor source of truth for AI provenance: rebases and squashes destroy attribution, humans may commit AI code, and AI agents may commit human-reviewed code. Source-level tags survive all of these operations.

## How to adopt this in a new repo

1. Copy [`examples/AI_DISCLOSURE.md`](./examples/AI_DISCLOSURE.md) into your repo root and customise the defaults.
2. Link to this convention from your README: *"AI disclosure: see [AI_DISCLOSURE.md](./AI_DISCLOSURE.md)."*
3. Add file-level headers to new files as you create or modify them. Do not try to retroactively annotate the entire repo at once.
4. Optionally, add an AI-disclosure field to your PR template: *"AI assistance in this PR: model, scope, level."*

## What this convention is not

This list is intended to head off the most common misreadings of the convention.

- Not a quality signal. `ai-generated` does not mean "lower quality". `none` does not mean "well written". The tag describes provenance only.
- Not enforceable. Adoption is voluntary. The point is transparency, not policing.
- Not language-specific. It works in any language with a comment syntax.
- Not git-based. Tags live in source files, not commit metadata, so they survive rebases, squashes, and rewrites.

## Acknowledgements

This convention is built on prior work by:

- The [W3C AI Content Disclosure Community Group](https://www.w3.org/community/ai-content-disclosure/) for the four-level vocabulary.
- The [SPDX Project](https://spdx.dev/) for the `SPDX-` line-tag format that this convention follows.
- [Anil Madhavapeddy](https://anil.recoil.org/notes/opam-ai-disclosure) for the OCaml-specific proposal that inspired this language-agnostic generalisation.
- [David E. Weekly](https://github.com/dweekly/ai-content-disclosure) for the W3C reference implementation.

## License

This specification is released under [CC0 1.0 Universal](./LICENSE). Fork it, modify it, or ignore it. No attribution required.
