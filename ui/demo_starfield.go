package ui

import (
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var starStyles = []lipgloss.Style{
	lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Bold(true), // Very close (bright white)
	lipgloss.NewStyle().Foreground(lipgloss.Color("#8be9fd")),            // Close (cyan)
	lipgloss.NewStyle().Foreground(lipgloss.Color("#f8f8f2")),            // Mid-range (white/gray)
	lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4")),            // Far (grey)
	lipgloss.NewStyle().Foreground(lipgloss.Color("#44475a")),            // Very far (dim grey)
}

// renderStarfield renders the 3D Starfield warp.
func (m Model) renderStarfield() string {
	w := m.TerminalWidth
	h := m.TerminalHeight
	if w <= 0 || h <= 0 {
		return ""
	}

	centerX := float64(w) / 2.0
	centerY := float64(h) / 2.0

	// Create grid
	grid := make([][]string, h)
	for y := 0; y < h; y++ {
		grid[y] = make([]string, w)
		for x := 0; x < w; x++ {
			grid[y][x] = " "
		}
	}

	for _, star := range m.Stars {
		if star.Z <= 0 {
			continue
		}

		// Perspective projection
		// typical character is twice as tall as it is wide, so compensate with a factor of ~1.8
		px := centerX + (star.X/star.Z)*1.8
		py := centerY + (star.Y/star.Z)

		ix := int(px)
		iy := int(py)

		if ix >= 0 && ix < w && iy >= 0 && iy < h {
			var char string
			var style lipgloss.Style

			if star.Z < 0.2 {
				char = "●"
				style = starStyles[0]
			} else if star.Z < 0.4 {
				char = "O"
				style = starStyles[1]
			} else if star.Z < 0.6 {
				char = "o"
				style = starStyles[2]
			} else if star.Z < 0.8 {
				char = "•"
				style = starStyles[3]
			} else {
				char = "·"
				style = starStyles[4]
			}

			grid[iy][ix] = style.Render(char)
		}
	}

	// Join rows
	rows := make([]string, h)
	for y := 0; y < h; y++ {
		rows[y] = strings.Join(grid[y], "")
	}
	return strings.Join(rows, "\n")
}

// updateStarfield updates the 3D Starfield simulation positions.
func (m Model) updateStarfield() (Model, tea.Cmd) {
	w := m.TerminalWidth
	h := m.TerminalHeight
	if w <= 0 || h <= 0 {
		return m, takeoverTick(time.Second / 30)
	}

	const maxStars = 100

	// Initialize stars if empty
	if len(m.Stars) == 0 {
		stars := make([]Star, maxStars)
		for i := 0; i < maxStars; i++ {
			stars[i] = Star{
				X: (rand.Float64()*2.0 - 1.0) * 80.0,
				Y: (rand.Float64()*2.0 - 1.0) * 40.0,
				Z: rand.Float64(),
			}
		}
		m.Stars = stars
	}

	// Update stars
	for i := 0; i < len(m.Stars); i++ {
		star := &m.Stars[i]
		star.Z -= 0.02 // Travel towards screen
		if star.Z <= 0.0 {
			// Reset star to back of field
			star.X = (rand.Float64()*2.0 - 1.0) * 80.0
			star.Y = (rand.Float64()*2.0 - 1.0) * 40.0
			star.Z = 1.0
		}
	}

	return m, takeoverTick(time.Second / 30)
}
