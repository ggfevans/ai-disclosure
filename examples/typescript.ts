/**
 * SPDX-AI-Disclosure: ai-generated
 * SPDX-AI-Model: claude-opus-4-6
 * SPDX-AI-Provider: Anthropic
 * SPDX-AI-Scope: Full module generated; type definitions reviewed.
 * SPDX-AI-Date: 2026-05-17
 */

export type DisclosureLevel =
  | 'none'
  | 'ai-assisted'
  | 'ai-generated'
  | 'autonomous';

export interface DisclosureTag {
  level: DisclosureLevel;
  model?: string;
  provider?: string;
  scope?: string;
  date?: string;
}

export function isDisclosed(tag: Partial<DisclosureTag>): tag is DisclosureTag {
  return typeof tag.level === 'string';
}
