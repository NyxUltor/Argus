//go:build ssh

package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func getHostModeString() string {
	return "ssh-secure / network"
}

func handleFallbackCommand(m Model, cmdText string) (Model, tea.Cmd) {
	m.History = append(m.History, HistoryItem{
		Command: cmdText,
		Output:  "argus: command not found: " + cmdText,
	})
	return m, nil
}
