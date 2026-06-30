package ui

// layout.go — pure layout/math helpers with no rendering side-effects.
// These functions are free of lipgloss styling, making them unit-testable
// independently of the full Model.

import "strings"

// wrapText hard-wraps text to width columns, respecting existing newlines
// and breaking overlong words at the width boundary.
func wrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		if len(line) == 0 {
			result = append(result, "")
			continue
		}
		words := strings.Fields(line)
		current := ""
		for _, word := range words {
			wordRunes := []rune(word)
			spaceNeeded := 0
			if current != "" {
				spaceNeeded = 1
			}
			if len([]rune(current))+spaceNeeded+len(wordRunes) > width {
				if current != "" {
					result = append(result, current)
					current = ""
				}
				for len(wordRunes) > width {
					result = append(result, string(wordRunes[:width]))
					wordRunes = wordRunes[width:]
				}
				current = string(wordRunes)
			} else {
				if current == "" {
					current = word
				} else {
					current += " " + word
				}
			}
		}
		if current != "" {
			result = append(result, current)
		}
	}
	return result
}

// MaxScrollOffset returns the maximum valid value for m.ScrollOffset.
// Returns 0 when history fits entirely within the viewport (no scrolling needed).
func (m Model) MaxScrollOffset() int {
	H_history := m.TerminalHeight - len(strings.Split(m.renderStatsPanel(), "\n")) - len(strings.Split(m.renderInputSection(), "\n")) - 2
	// Clamp H_history to minimum 1 to prevent zero/negative denominators in
	// scroll ratio math and invalid slice allocations on very small terminals.
	if H_history < 1 {
		H_history = 1
	}
	maxScroll := len(m.renderHistoryLines()) - H_history
	if maxScroll < 0 {
		return 0
	}
	return maxScroll
}

// logoOriginCol returns the column offset of the logo in the stats/header panel.
func (m Model) logoOriginCol() int {
	return 2
}

// headerHeight returns the number of rows occupied by the header panel.
func (m Model) headerHeight() int {
	return len(EyeArt)
}

