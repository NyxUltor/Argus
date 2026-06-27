package ui

import "time"

// TakeoverType identifies which fullscreen demo is running.
type TakeoverType string

const (
	TakeoverNone      TakeoverType = ""
	TakeoverMatrix    TakeoverType = "matrix"
	TakeoverStarfield TakeoverType = "starfield"
)

// Ball is a single physics object in overlay mode.
type Ball struct {
	X, Y   float64
	VX, VY float64
	Char   rune
	Color  string // lipgloss hex
}

// FloatingEyeState is the autonomous easter egg eye.
type FloatingEyeState struct {
	X, Y   float64
	VX, VY float64
	TTL    time.Duration
}
