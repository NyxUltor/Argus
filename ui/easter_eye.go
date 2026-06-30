package ui

import (
	"math"
	"math/rand"
	"sync"
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

var (
	logoHitmap [][]bool
	hitmapOnce sync.Once
)

func eyeTick() tea.Cmd {
	return tea.Tick(eyeTickRate, func(time.Time) tea.Msg { return eyeTickMsg{} })
}

func randomEyeCooldown() time.Duration {
	spread := eyeMaxCooldown - eyeMinCooldown
	return eyeMinCooldown + time.Duration(rand.Int63n(int64(spread)))
}

// floatingEyeSpriteLines returns plain (unstyled) braille lines for the easter egg eye.
func floatingEyeSpriteLines() []string {
	return []string{
		"⠀⣠⣾⣷⣄⠀",
		"⣼⣿⣯⣿⣿⣧",
		"⠙⢿⣿⣿⡿⠋",
		"⠀⠀⠙⠋⠀⠀",
	}
}

// floatingEyeWidth returns the display width of the sprite in cells.
func floatingEyeWidth() int {
	return lipgloss.Width(floatingEyeSpriteLines()[0])
}

func getLogoHitmap() [][]bool {
	hitmapOnce.Do(func() {
		logoHitmap = make([][]bool, len(EyeArt))
		for r, line := range EyeArt {
			runes := []rune(line)
			logoHitmap[r] = make([]bool, len(runes))
			for c, ch := range runes {
				// Exclude regular space ' ' and U+2800 (braille empty) '⠀'
				if ch != ' ' && ch != '\u2800' {
					logoHitmap[r][c] = true
				}
			}
		}
	})
	return logoHitmap
}

func clampVelocity(vx, vy float64, minMag, maxMag float64) (float64, float64) {
	mag := math.Sqrt(vx*vx + vy*vy)
	if mag == 0 {
		angle := rand.Float64() * 2 * math.Pi
		return math.Cos(angle) * minMag, math.Sin(angle) * minMag
	}
	if mag < minMag {
		scale := minMag / mag
		return vx * scale, vy * scale
	}
	if mag > maxMag {
		scale := maxMag / mag
		return vx * scale, vy * scale
	}
	return vx, vy
}

func overlapsLogo(x, y int, hitmap [][]bool, logoCol int) bool {
	r := y
	c := x - logoCol
	if r >= 0 && r < len(hitmap) {
		if c >= 0 && c < len(hitmap[r]) {
			return hitmap[r][c]
		}
	}
	return false
}

func (m Model) updateFloatingEye() Model {
	tick := eyeTickRate

	if m.FloatingEye == nil {
		m.EyeCooldownRemaining -= tick
		if m.EyeCooldownRemaining <= 0 {
			m = spawnFloatingEye(m)
		}
		return m
	}

	eye := m.FloatingEye
	eye.TTL -= tick

	// Brownian drift nudge:
	// Base drift nudge: ±0.15 range. Dazed drift nudge: ±0.3 range.
	var nudge float64
	if eye.DazedRemaining > 0 {
		nudge = 0.3
	} else {
		nudge = 0.15
	}
	eye.VX += (rand.Float64() * 2 * nudge) - nudge
	eye.VY += (rand.Float64() * 2 * nudge) - nudge

	// Speed clamp calculation
	var minSpeed, maxSpeed float64
	if eye.DazedRemaining > 0 {
		eye.DazedRemaining -= tick
		if eye.DazedRemaining <= 0 {
			eye.DazedRemaining = 0
			eye.Restabilizing = true
			eye.RestabilizeTimer = 0
		}
		minSpeed = 0.05
		maxSpeed = 0.4
	} else if eye.Restabilizing {
		eye.RestabilizeTimer += tick
		progress := float64(eye.RestabilizeTimer) / float64(500*time.Millisecond)
		if progress >= 1.0 {
			progress = 1.0
			eye.Restabilizing = false
			eye.RestabilizeTimer = 0
		}
		minSpeed = 0.05 + (0.2-0.05)*progress
		maxSpeed = 0.4 + (1.2-0.4)*progress
	} else {
		minSpeed = 0.2
		maxSpeed = 1.2
	}

	// Clamp speed vector
	eye.VX, eye.VY = clampVelocity(eye.VX, eye.VY, minSpeed, maxSpeed)

	// Move eye
	eye.X += eye.VX
	eye.Y += eye.VY

	// Collision detection
	hitmap := getLogoHitmap()
	logoCol := m.logoOriginCol()
	eyeW := floatingEyeWidth()
	eyeH := len(floatingEyeSpriteLines())

	tl := overlapsLogo(int(eye.X), int(eye.Y), hitmap, logoCol)
	tr := overlapsLogo(int(eye.X)+eyeW-1, int(eye.Y), hitmap, logoCol)
	bl := overlapsLogo(int(eye.X), int(eye.Y)+eyeH-1, hitmap, logoCol)
	br := overlapsLogo(int(eye.X)+eyeW-1, int(eye.Y)+eyeH-1, hitmap, logoCol)

	if tl || tr || bl || br {
		// Collision axis determination:
		// top/bottom -> reverse VY, left/right -> reverse VX, corner overlap -> reverse both
		reverseVX := false
		reverseVY := false

		if (tl && tr) && !(bl || br) {
			reverseVY = true
		} else if (bl && br) && !(tl || tr) {
			reverseVY = true
		} else if (tl && bl) && !(tr || br) {
			reverseVX = true
		} else if (tr && br) && !(tl || bl) {
			reverseVX = true
		} else {
			// Diagonal or corner overlap (e.g. tl only, tr only, or complex combination)
			reverseVX = true
			reverseVY = true
		}

		if reverseVX {
			eye.VX = -eye.VX
		}
		if reverseVY {
			eye.VY = -eye.VY
		}

		// Enter Dazed state
		eye.DazedRemaining = 1500 * time.Millisecond
		eye.Restabilizing = false
		eye.RestabilizeTimer = 0
	}

	// Despawn checks
	// "If any part of the eye sprite crosses the header boundary (any edge): despawn cleanly"
	offscreen := eye.X < 0 ||
		eye.X+float64(eyeW) > float64(m.TerminalWidth) ||
		eye.Y < 0 ||
		eye.Y+float64(eyeH) > float64(m.headerHeight())

	if eye.TTL <= 0 || offscreen {
		m.FloatingEye = nil
		m.EyeCooldownRemaining = randomEyeCooldown()
	}

	return m
}

func spawnFloatingEye(m Model) Model {
	var x, y, vx, vy float64
	edge := rand.Intn(3)
	eyeW := floatingEyeWidth()
	eyeH := len(floatingEyeSpriteLines())
	headerH := m.headerHeight()

	// Initial inward velocity magnitude: low (e.g. [0.3, 0.6])
	speed := 0.3 + rand.Float64()*0.3

	switch edge {
	case 0: // top edge
		minX := 2
		maxX := m.TerminalWidth - 2 - eyeW
		if maxX <= minX {
			x = float64(minX)
		} else {
			x = float64(minX + rand.Intn(maxX-minX+1))
		}
		y = 0
		vx = rand.Float64()*0.4 - 0.2
		vy = speed
	case 1: // bottom edge
		minX := 2
		maxX := m.TerminalWidth - 2 - eyeW
		if maxX <= minX {
			x = float64(minX)
		} else {
			x = float64(minX + rand.Intn(maxX-minX+1))
		}
		y = float64(headerH - eyeH)
		vx = rand.Float64()*0.4 - 0.2
		vy = -speed
	case 3: // right edge
		x = float64(m.TerminalWidth - eyeW)
		minY := 2
		maxY := headerH - 2 - eyeH
		if maxY <= minY {
			y = float64(minY)
		} else {
			y = float64(minY + rand.Intn(maxY-minY+1))
		}
		vx = -speed
		vy = rand.Float64()*0.4 - 0.2
	}

	ttlSpread := eyeMaxTTL - eyeMinTTL
	m.FloatingEye = &FloatingEyeState{
		X:                x,
		Y:                y,
		VX:               vx,
		VY:               vy,
		TTL:              eyeMinTTL + time.Duration(rand.Int63n(int64(ttlSpread))),
		DazedRemaining:   0,
		Restabilizing:    false,
		RestabilizeTimer: 0,
	}
	return m
}
