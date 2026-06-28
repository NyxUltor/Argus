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
	// Takeover mode bypasses the normal layout entirely.
	if m.ActiveTakeover != TakeoverNone {
		return m.renderTakeover()
	}
	if m.TerminalWidth == 0 || m.TerminalHeight == 0 {
		return "Initializing screen buffer..."
	}

	scr := newScreen(m.TerminalHeight, m.TerminalWidth)

	// 1. Header panel (top)
	headerLines := m.headerLines()
	for i, line := range headerLines {
		m.writePlainLine(scr, i, line)
	}
	H_top := len(headerLines)

	// 2. Input section (bottom)
	inputLines := m.inputLines()
	H_input := len(inputLines)
	inputStartRow := m.TerminalHeight - H_input
	for i, line := range inputLines {
		m.writePlainLine(scr, inputStartRow+i, line)
	}

	// 3. Divider borders
	borderRow := H_top       // purple border under header (═══)
	dividerRow := inputStartRow - 1 // grey divider above input (───)
	m.writeBorderLine(scr, borderRow, '═', borderStyle1)
	m.writeBorderLine(scr, dividerRow, '─', borderStyle2)

	// 4. History viewport
	H_history := dividerRow - borderRow - 1
	if H_history < 1 {
		H_history = 1
	}
	historyStartRow := borderRow + 1
	m.writeHistorySection(scr, historyStartRow, H_history)

	// 5. Stamp sprites on top
	for _, sp := range m.buildSprites() {
		scr.stampSprite(sp.Row, sp.Col, sp.Lines, sp.Style)
	}

	return scr.emit()
}

// writePlainLine parses a styled string by calling lipgloss.Width and writes
// it into the screen. Because we can't trivially decompose a lipgloss-rendered
// string back into (rune, style) pairs, we use a different approach:
// write the FULL rendered string as a single-cell span at column 0 using
// the zero style, letting emit() output it verbatim.
//
// For lines that don't need sprites over them (header, input, borders),
// this is fine — sprites only appear in the history region.
//
// Implementation: store the entire rendered line as a "raw ANSI" span.
// We extend Cell to support a RawANSI mode for this purpose.
func (m Model) writePlainLine(scr Screen, row int, rendered string) {
	if row < 0 || row >= len(scr) {
		return
	}
	// Store the entire rendered string in col 0 as a raw cell.
	// All other cells in the row remain spaces.
	scr[row][0] = Cell{Ch: 0, Raw: rendered}
}

// writeBorderLine fills a row with a repeated border character using the given style.
func (m Model) writeBorderLine(scr Screen, row int, ch rune, style lipgloss.Style) {
	if row < 0 || row >= len(scr) {
		return
	}
	for col := 0; col < len(scr[row]); col++ {
		scr[row][col] = Cell{Ch: ch, Style: style}
	}
}

// writeHistorySection renders the scrollable history panel into the screen,
// including the scrollbar. Only rows in [startRow, startRow+height) are written.
func (m Model) writeHistorySection(scr Screen, startRow, height int) {
	historyLines := m.renderHistoryLines()

	var visibleLines []string
	showScrollbar := false
	handlePos := 0

	if len(historyLines) > height {
		showScrollbar = true
		start := len(historyLines) - height - m.ScrollOffset
		end := len(historyLines) - m.ScrollOffset
		if start < 0 {
			start = 0
		}
		if end > len(historyLines) {
			end = len(historyLines)
		}
		visibleLines = historyLines[start:end]

		maxScroll := len(historyLines) - height
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
		if height > 1 {
			handlePos = (height - 1) - int(ratio*float64(height-1)+0.5)
		}
	} else {
		// Pad top with blank lines to pin content to the bottom.
		padding := height - len(historyLines)
		visibleLines = append(make([]string, padding), historyLines...)
	}

	for i := 0; i < height; i++ {
		row := startRow + i
		if i < len(visibleLines) {
			m.writePlainLine(scr, row, visibleLines[i])
		}
		// Scrollbar in the last two columns.
		if showScrollbar {
			if i == handlePos {
				scr.set(row, m.TerminalWidth-1, Cell{Ch: '█', Style: borderStyle1})
			} else {
				scr.set(row, m.TerminalWidth-1, Cell{Ch: '│', Style: borderStyle2})
			}
		}
	}
}

// headerLines returns the rendered header panel as a slice of ANSI strings.
func (m Model) headerLines() []string {
	return strings.Split(m.renderStatsPanel(), "\n")
}

// inputLines returns the rendered input section as a slice of ANSI strings.
func (m Model) inputLines() []string {
	return strings.Split(m.renderInputSection(), "\n")
}
