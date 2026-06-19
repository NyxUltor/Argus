package ui

import "fmt"

// View renders the layout to the terminal screen.
func (m Model) View() string {
	return fmt.Sprintf(
		"Welcome to NYX Terminal Portfolio\n"+
			"Resolution: %dx%d\n\n"+
			"Press 'q' or 'Ctrl+C' to exit local testing.",
		m.TerminalWidth, m.TerminalHeight,
	)
}
