package ui

import tea "github.com/charmbracelet/bubbletea"

// Init initializes the terminal session state.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles incoming key events and terminal size changes.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Our safe local test exit keys
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height
	}
	return m, nil
}
