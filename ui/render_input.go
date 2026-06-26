package ui

// render_input.go — renders the persistent bottom input section:
// suggestions overlay, prompt line, and cursor.

import (
	"fmt"
	"strings"
)

// renderSuggestions renders the autocomplete hints line above the prompt.
// Returns an empty string when there are no suggestions (no line is emitted).
func (m Model) renderSuggestions() string {
	if len(m.Suggestions) == 0 {
		return ""
	}
	// idea for later: allow navigating/selecting suggestions with arrow keys
	return borderStyle2.Italic(true).Render(
		fmt.Sprintf("  (Suggestions: %s)", strings.Join(m.Suggestions, ", ")),
	)
}

// renderInputLine renders the active prompt with a block cursor at CursorPos.
func (m Model) renderInputLine() string {
	prompt := promptStyle.Render(fmt.Sprintf("[%s@ARGUS ~]$ ", m.SystemUsername))
	runes := []rune(m.InputBuffer)

	cursorPos := m.CursorPos
	if cursorPos < 0 {
		cursorPos = 0
	}
	if cursorPos > len(runes) {
		cursorPos = len(runes)
	}

	var styledInput string
	if cursorPos == len(runes) {
		// Cursor is past the last character — show a trailing block.
		styledInput = string(runes) + cursorStyle.Render(" ")
	} else {
		before := string(runes[:cursorPos])
		charAtCursor := string(runes[cursorPos])
		after := string(runes[cursorPos+1:])
		styledInput = before + cursorStyle.Render(charAtCursor) + after
	}

	return prompt + styledInput
}

// renderInputSection combines the suggestions line (if present) and the prompt.
func (m Model) renderInputSection() string {
	var lines []string
	if sug := m.renderSuggestions(); sug != "" {
		lines = append(lines, sug)
	}
	lines = append(lines, m.renderInputLine())
	return strings.Join(lines, "\n")
}
