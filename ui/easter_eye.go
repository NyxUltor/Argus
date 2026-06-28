package ui

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	eyeMinCooldown = 30 * time.Second
	eyeMaxCooldown = 60 * time.Second
	eyeMinTTL      = 8 * time.Second
	eyeMaxTTL      = 15 * time.Second
	eyeTickRate    = time.Second / 20
)

func eyeTick() tea.Cmd {
	return tea.Tick(eyeTickRate, func(time.Time) tea.Msg { return eyeTickMsg{} })
}

func randomEyeCooldown() time.Duration {
	spread := eyeMaxCooldown - eyeMinCooldown
	return eyeMinCooldown + time.Duration(rand.Int63n(int64(spread)))
}

// floatingEyeSprite returns the pre-styled lines for the easter egg eye.
// TODO: replace placeholder with real braille eye art authored from reference images.
// floatingEyeSpriteLines returns plain (unstyled) braille lines for the easter egg eye.
// Style is applied by the compositor. TODO: replace with real art.
func floatingEyeSpriteLines() []string {
	return []string{
		"⠀⣠⣾⣷⣄⠀",
		"⣼⣿⣯⣿⣿⣧",
		"⠙⢿⣿⣿⡿⠋",
		"⠀⠀⠙⠛⠋⠀⠀",
	}
}

// floatingEyeWidth returns the display width of the sprite in cells.
func floatingEyeWidth() int {
	return lipgloss.Width(floatingEyeSpriteLines()[0])
}

func (m Model) updateFloatingEye() Model {
	tick := eyeTickRate

	if m.FloatingEye == nil {
		// Count down cooldown
		m.EyeCooldownRemaining -= tick
		if m.EyeCooldownRemaining <= 0 {
			m = spawnFloatingEye(m)
		}
		return m
	}

	// Move eye
	m.FloatingEye.X += m.FloatingEye.VX
	m.FloatingEye.Y += m.FloatingEye.VY
	m.FloatingEye.TTL -= tick

	// Slight random drift perturbation
	m.FloatingEye.VX += (rand.Float64()*0.2 - 0.1)
	m.FloatingEye.VY += (rand.Float64()*0.2 - 0.1)

	// Clamp drift speed
	m.FloatingEye.VX = clamp(m.FloatingEye.VX, -1.5, 1.5)
	m.FloatingEye.VY = clamp(m.FloatingEye.VY, -0.8, 0.8)

	// Despawn conditions: TTL expired or reached an edge
	eyeW := float64(floatingEyeWidth())
	eyeH := float64(len(floatingEyeSpriteLines()))
	offscreen := m.FloatingEye.X < -eyeW ||
		m.FloatingEye.X > float64(m.TerminalWidth) ||
		m.FloatingEye.Y < -eyeH ||
		m.FloatingEye.Y > float64(m.TerminalHeight)

	if m.FloatingEye.TTL <= 0 || offscreen {
		m.FloatingEye = nil
		m.EyeCooldownRemaining = randomEyeCooldown()
	}

	return m
}

func spawnFloatingEye(m Model) Model {
	// Spawn at a random position along any of the 4 edges
	var x, y, vx, vy float64
	edge := rand.Intn(4)
	speed := 0.4 + rand.Float64()*0.6
	switch edge {
	case 0: // top
		x = rand.Float64() * float64(m.TerminalWidth)
		y = 0
		vx = rand.Float64()*1.0 - 0.5
		vy = speed
	case 1: // bottom
		x = rand.Float64() * float64(m.TerminalWidth)
		y = float64(m.TerminalHeight - 1)
		vx = rand.Float64()*1.0 - 0.5
		vy = -speed
	case 2: // left
		x = 0
		y = rand.Float64() * float64(m.TerminalHeight)
		vx = speed
		vy = rand.Float64()*0.6 - 0.3
	case 3: // right
		x = float64(m.TerminalWidth - 1)
		y = rand.Float64() * float64(m.TerminalHeight)
		vx = -speed
		vy = rand.Float64()*0.6 - 0.3
	}

	ttlSpread := eyeMaxTTL - eyeMinTTL
	m.FloatingEye = &FloatingEyeState{
		X:   x,
		Y:   y,
		VX:  vx,
		VY:  vy,
		TTL: eyeMinTTL + time.Duration(rand.Int63n(int64(ttlSpread))),
	}
	return m
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
