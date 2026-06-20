//go:build !ssh

package ui

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// getHostModeString returns the status indicator text for local environments
func getHostModeString() string {
	return "local / sh-exec"
}

// handleFallbackCommand forks unhandled shell inputs directly over to bash asynchronously
func handleFallbackCommand(m Model, cmdText string) (Model, tea.Cmd) {
	m.History = append(m.History, HistoryItem{
		Command: cmdText,
		Output:  "Running...",
	})
	cmdIndex := len(m.History) - 1

	return m, func() tea.Msg {
		cmd := exec.Command("bash", "-c", cmdText)
		output, err := cmd.CombinedOutput()
		outStr := string(output)
		if err != nil && outStr == "" {
			outStr = err.Error()
		}
		return commandFinishedMsg{
			Index:  cmdIndex,
			Output: outStr,
		}
	}
}
