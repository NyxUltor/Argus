package ui

import (
	"math"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	BallCap      = 20
	ballTickRate = time.Second / 30
	ballGravity  = 0.4
	ballDamping  = 0.75
	ballMinSpeed = 0.3
)

var ballChars = []rune{'O', '●', '◉', 'o', '0'}
var ballColors = []string{"#ff79c6", "#bd93f9", "#50fa7b", "#ffb86c", "#ff5555", "#8be9fd", "#f1fa8c"}

func ballTick() tea.Cmd {
	return tea.Tick(ballTickRate, func(time.Time) tea.Msg { return ballTickMsg{} })
}

func spawnBalls(m Model, n int) Model {
	if n < 0 {
		n = -n
	}
	// Note: user commented: "-0 just means dont summon anything no flag defaults to 1".
	// Therefore, if n is explicitly 0, we do not override it to 1; we spawn 0 balls.
	canSpawn := BallCap - len(m.Balls)
	if n > canSpawn {
		n = canSpawn
	}
	for i := 0; i < n; i++ {
		m.Balls = append(m.Balls, Ball{
			X:     float64(5 + rand.Intn(max(m.TerminalWidth-10, 1))),
			Y:     1,
			VX:    (rand.Float64()*4 - 2),
			VY:    rand.Float64()*1.5 + 0.5,
			Char:  ballChars[rand.Intn(len(ballChars))],
			Color: ballColors[rand.Intn(len(ballColors))],
		})
	}
	return m
}

// ballFloor returns the Y row of the border line above the input section.
func (m Model) ballFloor() float64 {
	inputH := len(strings.Split(m.renderInputSection(), "\n"))
	return float64(m.TerminalHeight - inputH - 2)
}

func (m Model) updateBalls() Model {
	floor := m.ballFloor()
	alive := m.Balls[:0]

	for i := range m.Balls {
		b := &m.Balls[i]
		b.VY += ballGravity
		b.X += b.VX
		b.Y += b.VY

		// Wall collisions
		if b.X < 0 {
			b.X = 0
			b.VX = math.Abs(b.VX) * ballDamping
		}
		if b.X >= float64(m.TerminalWidth-1) {
			b.X = float64(m.TerminalWidth - 2)
			b.VX = -math.Abs(b.VX) * ballDamping
		}
		if b.Y < 0 {
			b.Y = 0
			b.VY = math.Abs(b.VY) * ballDamping
		}
		if b.Y >= floor {
			b.Y = floor
			b.VY = -math.Abs(b.VY) * ballDamping
			b.VX *= ballDamping
		}

		// Easter egg eye collision
		if m.FloatingEye != nil {
			eyeW := float64(floatingEyeWidth())
			eyeH := float64(len(floatingEyeSpriteLines()))
			if b.X >= m.FloatingEye.X && b.X <= m.FloatingEye.X+eyeW &&
				b.Y >= m.FloatingEye.Y && b.Y <= m.FloatingEye.Y+eyeH {
				if b.X < m.FloatingEye.X+eyeW/2 {
					b.VX = -math.Abs(b.VX)
				} else {
					b.VX = math.Abs(b.VX)
				}
				if b.Y < m.FloatingEye.Y+eyeH/2 {
					b.VY = -math.Abs(b.VY)
				} else {
					b.VY = math.Abs(b.VY)
				}
			}
		}

		// Despawn
		if math.Abs(b.VX) < ballMinSpeed && math.Abs(b.VY) < ballMinSpeed && b.Y >= floor-0.5 {
			continue
		}
		alive = append(alive, *b)
	}

	// Ball-to-ball collisions
	for i := range alive {
		for j := i + 1; j < len(alive); j++ {
			dx := alive[i].X - alive[j].X
			dy := alive[i].Y - alive[j].Y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < 1.5 {
				alive[i].VX, alive[j].VX = alive[j].VX*ballDamping, alive[i].VX*ballDamping
				alive[i].VY, alive[j].VY = alive[j].VY*ballDamping, alive[i].VY*ballDamping
			}
		}
	}

	m.Balls = alive
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
