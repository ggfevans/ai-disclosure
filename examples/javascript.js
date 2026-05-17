/**
 * SPDX-AI-Disclosure: ai-assisted
 * SPDX-AI-Model: claude-opus-4-6
 * SPDX-AI-Provider: Anthropic
 * SPDX-AI-Scope: Helper function generated; integration written by hand.
 * SPDX-AI-Date: 2026-05-17
 */

export function slugify(input) {
  return input
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '');
}
