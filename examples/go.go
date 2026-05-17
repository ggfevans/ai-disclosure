// SPDX-AI-Disclosure: ai-generated
// SPDX-AI-Model: claude-opus-4-6
// SPDX-AI-Provider: Anthropic
// SPDX-AI-Scope: Parser logic generated; struct fields by hand.
// SPDX-AI-Date: 2026-05-17

package disclosure

import (
	"regexp"
)

// Disclosure represents the parsed SPDX-AI-* tags from a source file.
type Disclosure struct {
	Level    string
	Model    string
	Provider string
	Scope    string
	Date     string
}

var tagPattern = regexp.MustCompile(`(?m)SPDX-AI-(\w+):\s*(.+?)\s*$`)

// Parse extracts disclosure tags from the source string. Returns nil if no
// SPDX-AI-Disclosure tag is found.
func Parse(source string) *Disclosure {
	matches := tagPattern.FindAllStringSubmatch(source, -1)
	if len(matches) == 0 {
		return nil
	}
	fields := make(map[string]string)
	for _, m := range matches {
		fields[m[1]] = m[2]
	}
	level, ok := fields["Disclosure"]
	if !ok {
		return nil
	}
	return &Disclosure{
		Level:    level,
		Model:    fields["Model"],
		Provider: fields["Provider"],
		Scope:    fields["Scope"],
		Date:     fields["Date"],
	}
}
