package ui

import (
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) handleSummon(cmdClean string) (tea.Model, tea.Cmd) {
	parts := strings.Fields(cmdClean) // ["summon", "ball", "-5"]
	if len(parts) < 2 {
		m.History = append(m.History, HistoryItem{
			Command: cmdClean,
			Output:  "Usage: summon <target> [args]\nType 'fun' for available targets.",
		})
		return m, nil
	}

	switch parts[1] {
	case "ball":
		n := 1
		if len(parts) >= 3 {
			raw := parts[2]
			// Strip leading dash(es) and parse
			raw = strings.TrimLeft(raw, "-")
			if v, err := strconv.Atoi(raw); err == nil {
				n = v
			}
		}
		m = spawnBalls(m, n)
		var cmd tea.Cmd
		if len(m.Balls) > 0 {
			cmd = ballTick()
		}
		return m, cmd

	case "matrix":
		m.Balls = []Ball{}
		m.ActiveTakeover = TakeoverMatrix
		return m, takeoverTick(time.Second / 15)

	case "starfield":
		m.Balls = []Ball{}
		m.ActiveTakeover = TakeoverStarfield
		return m, takeoverTick(time.Second / 24)

	default:
		m.History = append(m.History, HistoryItem{
			Command: cmdClean,
			Output:  "Unknown summon target: " + parts[1] + "\nType 'fun' for available targets.",
		})
		return m, nil
	}
}
