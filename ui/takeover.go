package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func takeoverTick(rate time.Duration) tea.Cmd {
	return tea.Tick(rate, func(time.Time) tea.Msg { return takeoverTickMsg{} })
}

// renderTakeover dispatches to the active demo's renderer.
// Returns a full-screen string replacing the normal view entirely.
func (m Model) renderTakeover() string {
	switch m.ActiveTakeover {
	case TakeoverMatrix:
		return m.renderMatrix()
	case TakeoverStarfield:
		return m.renderStarfield()
	default:
		return ""
	}
}

// clearTakeover stops any active takeover and returns to portfolio view.
func (m Model) clearTakeover() Model {
	m.ActiveTakeover = TakeoverNone
	return m
}
