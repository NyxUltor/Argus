package ui

// render_history.go — converts m.History into display-ready string lines.
// This file is purely about translating model state into visual rows;
// scroll windowing and scrollbar rendering happen in view.go.

import "fmt"

func (m Model) renderHistoryLines() []string {
	var lines []string
	for _, item := range m.History {
		// Each history entry begins with the prompt line showing the original command.
		promptLine := promptStyle.Render(fmt.Sprintf("[%s@ARGUS ~]$ ", m.SystemUsername)) + item.Command
		lines = append(lines, promptLine)

		// Wrap and append command output, if any.
		if item.Output != "" {
			outputLines := wrapText(item.Output, m.TerminalWidth-2)
			lines = append(lines, outputLines...)
		}
	}
	return lines
}
