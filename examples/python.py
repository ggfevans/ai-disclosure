# SPDX-AI-Disclosure: ai-generated
# SPDX-AI-Model: claude-opus-4-6
# SPDX-AI-Provider: Anthropic
# SPDX-AI-Scope: Implementation only; tests hand-written.
# SPDX-AI-Date: 2026-05-17

"""Parse SPDX-AI-Disclosure tags from source files."""

import re
from dataclasses import dataclass
from typing import Optional


TAG_PATTERN = re.compile(r"SPDX-AI-(\w+):\s*(.+?)\s*$", re.MULTILINE)


@dataclass
class Disclosure:
    level: str
    model: Optional[str] = None
    provider: Optional[str] = None
    scope: Optional[str] = None
    date: Optional[str] = None


def parse(source: str) -> Optional[Disclosure]:
    """Return a Disclosure if the source contains SPDX-AI tags, else None."""
    matches = dict(TAG_PATTERN.findall(source))
    if "Disclosure" not in matches:
        return None
    return Disclosure(
        level=matches["Disclosure"],
        model=matches.get("Model"),
        provider=matches.get("Provider"),
        scope=matches.get("Scope"),
        date=matches.get("Date"),
    )
