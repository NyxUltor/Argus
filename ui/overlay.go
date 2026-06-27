package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Sprite is a positioned renderable stamped over the base view.
type Sprite struct {
	X, Y  int
	Lines []string // pre-styled via lipgloss
}

// stampOverlay composites sprites onto a base view string.
// It splits the base into a rune grid, overwrites cells at each sprite's
// position (bounds-checked), then rejoins into a string.
// ANSI codes in base are not stripped — sprites are stamped as plain
// rune replacements which may interact with existing ANSI sequences.
// Keep sprite content plain or pre-styled and terminated with a reset.
func stampOverlay(base string, sprites []Sprite) string {
	rows := strings.Split(base, "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}

	for _, sp := range sprites {
		for dy, line := range sp.Lines {
			ry := sp.Y + dy
			if ry < 0 || ry >= len(grid) {
				continue
			}
			for dx, ch := range []rune(line) {
				rx := sp.X + dx
				if rx < 0 || rx >= len(grid[ry]) {
					continue
				}
				grid[ry][rx] = ch
			}
		}
	}

	result := make([]string, len(grid))
	for i, row := range grid {
		result[i] = string(row)
	}
	return strings.Join(result, "\n")
}

// buildSprites collects all active overlay sprites from the model.
func (m Model) buildSprites() []Sprite {
	var sprites []Sprite
	for _, b := range m.Balls {
		sprites = append(sprites, Sprite{
			X:     int(b.X),
			Y:     int(b.Y),
			Lines: []string{lipgloss.NewStyle().Foreground(lipgloss.Color(b.Color)).Render(string(b.Char))},
		})
	}
	if m.FloatingEye != nil {
		sprites = append(sprites, Sprite{
			X:     int(m.FloatingEye.X),
			Y:     int(m.FloatingEye.Y),
			Lines: floatingEyeSprite(),
		})
	}
	return sprites
}
