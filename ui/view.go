package ui

// view.go — thin orchestrator. Assembles pre-rendered sections into the
// final full-screen layout with scroll windowing and the scrollbar track.
// Rendering logic lives in render_header.go, render_history.go, render_input.go.
// Layout math lives in layout.go. Styles live in styles.go.

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.TerminalWidth == 0 || m.TerminalHeight == 0 {
		return "Initializing screen buffer..."
	}

	// 1. Build the top stats panel and measure its height.
	topSection := m.renderStatsPanel()
	H_top := len(strings.Split(topSection, "\n"))

	// 2. Build the bottom input section and measure its height.
	inputSection := m.renderInputSection()
	H_input := len(strings.Split(inputSection, "\n"))

	// 3. Calculate the scrollable history viewport height.
	// Minimum of 1 to prevent zero-size allocations on very small terminals.
	H_history := m.TerminalHeight - H_top - H_input - 2
	if H_history < 1 {
		H_history = 1
	}

	historyLines := m.renderHistoryLines()

	var historySection string
	if len(historyLines) > H_history {
		// --- Scrolling path: content overflows the viewport ---
		start := len(historyLines) - H_history - m.ScrollOffset
		end := len(historyLines) - m.ScrollOffset
		if start < 0 {
			start = 0
		}
		if end > len(historyLines) {
			end = len(historyLines)
		}
		visibleLines := historyLines[start:end]

		// Scrollbar ratio — guarded against zero-division (Bug 2).
		maxScroll := len(historyLines) - H_history
		ratio := 0.0
		if maxScroll > 0 {
			ratio = float64(m.ScrollOffset) / float64(maxScroll)
		}
		if ratio < 0 {
			ratio = 0
		}
		if ratio > 1 {
			ratio = 1
		}

		handlePos := 0
		if H_history > 1 {
			handlePos = (H_history - 1) - int(ratio*float64(H_history-1)+0.5)
		}

		finalLines := make([]string, H_history)
		for i := 0; i < H_history; i++ {
			if i < len(visibleLines) {
				finalLines[i] = visibleLines[i]
			} else {
				finalLines[i] = strings.Repeat(" ", m.TerminalWidth-2)
			}
			// Pad the line to the full width so the scrollbar column is flush.
			// lipgloss.Width strips ANSI codes internally, giving the true cell count.
			visWidth := lipgloss.Width(finalLines[i])
			if padLen := (m.TerminalWidth - 2) - visWidth; padLen > 0 {
				finalLines[i] += strings.Repeat(" ", padLen)
			}
			// Append scrollbar glyph.
			if i == handlePos {
				finalLines[i] += " " + borderStyle1.Render("█")
			} else {
				finalLines[i] += " " + borderStyle2.Render("│")
			}
		}
		historySection = strings.Join(finalLines, "\n")

	} else {
		// --- No-scroll path: pad blank lines above to keep content pinned bottom ---
		paddingCount := H_history - len(historyLines)
		padding := make([]string, paddingCount)
		allLines := append(padding, historyLines...)
		historySection = strings.Join(allLines, "\n")
	}

	var s strings.Builder
	s.WriteString(topSection)
	s.WriteString("\n")
	s.WriteString(historySection)
	s.WriteString("\n")
	s.WriteString(borderStyle2.Render(strings.Repeat("─", m.TerminalWidth)))
	s.WriteString("\n")
	s.WriteString(inputSection)

	return s.String()
}
