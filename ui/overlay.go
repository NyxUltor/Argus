package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Cell is a single terminal cell: one display rune and its lipgloss style.
// The zero value renders as a space with no styling.
type Cell struct {
	Ch    rune
	Style lipgloss.Style
	Raw   string // if non-empty, emit Raw verbatim and skip Ch/Style
}

// Screen is a 2D grid of Cells representing the full terminal view.
// Row 0 is the top of the screen. Access: Screen[row][col].
type Screen [][]Cell

// newScreen allocates a Screen of the given dimensions filled with space cells.
func newScreen(rows, cols int) Screen {
	s := make(Screen, rows)
	for i := range s {
		s[i] = make([]Cell, cols)
		for j := range s[i] {
			s[i][j] = Cell{Ch: ' '}
		}
	}
	return s
}

// set writes a cell at (row, col), bounds-checking silently.
func (s Screen) set(row, col int, c Cell) {
	if row < 0 || row >= len(s) {
		return
	}
	if col < 0 || col >= len(s[row]) {
		return
	}
	s[row][col] = c
}

// writeString writes a plain string into the screen at (row, startCol)
// using the given style for every character. ANSI in the string is written
// literally as runes — do NOT pass pre-styled lipgloss output here.
// Pass plain text + style separately.
func (s Screen) writeString(row, startCol int, text string, style lipgloss.Style) {
	col := startCol
	for _, ch := range text {
		s.set(row, col, Cell{Ch: ch, Style: style})
		col++
	}
}

// stampSprite writes a sprite's lines into the screen at position (row, col).
// Each line in lines is a plain string (no ANSI). style is applied to every cell.
// For multi-style sprites, call stampSprite multiple times or use set() directly.
func (s Screen) stampSprite(row, col int, lines []string, style lipgloss.Style) {
	for dy, line := range lines {
		r := row + dy
		c := col
		for _, ch := range line {
			s.set(r, c, Cell{Ch: ch, Style: style})
			c++
		}
	}
}

// emit converts the Screen into a single ANSI string for Bubble Tea to render.
// Adjacent cells sharing the same style are grouped into one lipgloss Render call
// to minimise escape sequence volume.
func (s Screen) emit() string {
	var sb strings.Builder
	for r, row := range s {
		if r > 0 {
			sb.WriteByte('\n')
		}
		i := 0
		for i < len(row) {
			// Raw cell: emit verbatim, skip to next cell.
			if row[i].Raw != "" {
				sb.WriteString(row[i].Raw)
				i++
				continue
			}
			// Group consecutive non-raw cells with the same Style.
			j := i + 1
			for j < len(row) && row[j].Raw == "" && row[j].Style.String() == row[i].Style.String() {
				j++
			}
			var chunk strings.Builder
			for _, cell := range row[i:j] {
				chunk.WriteRune(cell.Ch)
			}
			sb.WriteString(row[i].Style.Render(chunk.String()))
			i = j
		}
	}
	return sb.String()
}

// Sprite is a positioned overlay element.
// Lines must be plain strings (no ANSI). Style is applied uniformly.
type Sprite struct {
	Row, Col int
	Lines    []string
	Style    lipgloss.Style
}

// buildSprites collects all active overlay sprites from the model.
func (m Model) buildSprites() []Sprite {
	var sprites []Sprite
	for _, b := range m.Balls {
		sprites = append(sprites, Sprite{
			Row:   int(b.Y),
			Col:   int(b.X),
			Lines: []string{string(b.Char)},
			Style: lipgloss.NewStyle().Foreground(lipgloss.Color(b.Color)),
		})
	}
	if m.FloatingEye != nil {
		sprites = append(sprites, Sprite{
			Row:   int(m.FloatingEye.Y),
			Col:   int(m.FloatingEye.X),
			Lines: floatingEyeSpriteLines(),
			Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#00ffd7")),
		})
	}
	return sprites
}
