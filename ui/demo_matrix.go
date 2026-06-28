package ui

import (
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var matrixColors = []lipgloss.Style{
	lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Bold(true), // Head (pure white)
	lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true), // Very bright green
	lipgloss.NewStyle().Foreground(lipgloss.Color("#30af5b")),            // Medium green
	lipgloss.NewStyle().Foreground(lipgloss.Color("#1a5f31")),            // Dim green
	lipgloss.NewStyle().Foreground(lipgloss.Color("#0d3018")),            // Very dim green
}

// renderMatrix renders the Matrix rain.
func (m Model) renderMatrix() string {
	w := m.TerminalWidth
	h := m.TerminalHeight
	if w <= 0 || h <= 0 {
		return ""
	}

	// Create a 2D grid of styled characters
	grid := make([][]string, h)
	for y := 0; y < h; y++ {
		grid[y] = make([]string, w)
		for x := 0; x < w; x++ {
			grid[y][x] = " "
		}
	}

	// Fill the grid based on column state
	for x := 0; x < w && x < len(m.MatrixColumns); x++ {
		col := m.MatrixColumns[x]
		leadY := int(col.Y)
		for dy := 0; dy < col.Length; dy++ {
			y := leadY - dy
			if y >= 0 && y < h {
				var char rune
				if len(col.Runes) > 0 {
					char = col.Runes[dy%len(col.Runes)]
				} else {
					char = ' '
				}

				if char != ' ' {
					// Choose styling based on distance from leadY
					var style lipgloss.Style
					if dy == 0 {
						style = matrixColors[0] // Head
					} else if dy < 3 {
						style = matrixColors[1] // Bright
					} else if dy < col.Length/2 {
						style = matrixColors[2] // Mid
					} else if dy < col.Length-2 {
						style = matrixColors[3] // Dim
					} else {
						style = matrixColors[4] // Fading
					}
					grid[y][x] = style.Render(string(char))
				}
			}
		}
	}

	// Join rows
	rows := make([]string, h)
	for y := 0; y < h; y++ {
		rows[y] = strings.Join(grid[y], "")
	}
	return strings.Join(rows, "\n")
}

// updateMatrix updates the Matrix rain column states and generates next tick.
func (m Model) updateMatrix() (Model, tea.Cmd) {
	w := m.TerminalWidth
	h := m.TerminalHeight
	if w <= 0 || h <= 0 {
		return m, takeoverTick(time.Second / 20)
	}

	// Initialize if size mismatch (resize or startup)
	if len(m.MatrixColumns) != w {
		cols := make([]MatrixColumn, w)
		for x := 0; x < w; x++ {
			cols[x] = MatrixColumn{
				Y:      float64(rand.Intn(max(h, 1))),
				Speed:  rand.Float64()*0.4 + 0.2,
				Length: rand.Intn(15) + 5,
				Runes:  generateRandomRunes(25),
			}
		}
		m.MatrixColumns = cols
	}

	// Move drops down
	for x := 0; x < w; x++ {
		col := &m.MatrixColumns[x]
		col.Y += col.Speed
		if int(col.Y)-col.Length >= h {
			col.Y = 0
			col.Speed = rand.Float64()*0.4 + 0.2
			col.Length = rand.Intn(15) + 5
			col.Runes = generateRandomRunes(col.Length + 5)
		} else {
			// Flicker / change runes slightly over time
			if rand.Float64() < 0.05 && len(col.Runes) > 0 {
				idx := rand.Intn(len(col.Runes))
				col.Runes[idx] = rune(33 + rand.Intn(93)) // Printable ASCII
			}
		}
	}

	return m, takeoverTick(time.Second / 20)
}

func generateRandomRunes(n int) []rune {
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		runes[i] = rune(33 + rand.Intn(93))
	}
	return runes
}
