package ui

// model.go — Model struct definition and constructor.
// Message types live in messages.go.

import (
	"os/user"
	"time"
)

// HistoryItem stores a single command and its output in the session history.
type HistoryItem struct {
	Command string
	Output  string
}

// Model is the central Bubble Tea application state.
type Model struct {
	InputBuffer    string
	CursorPos      int
	Suggestions    []string
	SelectedIndex  int
	ShowDropdown   bool
	ActiveView     string
	SelectedProj   string
	Message        string
	TerminalWidth  int
	TerminalHeight int
	SystemUsername string

	History      []HistoryItem
	ScrollOffset int

	// Animation — overlay
	Balls []Ball

	// Animation — takeover
	ActiveTakeover TakeoverType

	// Easter egg
	FloatingEye          *FloatingEyeState
	EyeCooldownRemaining time.Duration

	// Input tracking for iris
	MouseX int
	MouseY int

	// Matrix & Starfield animation states
	MatrixColumns []MatrixColumn
	Stars         []Star
}

// InitialModel constructs the default application state.
func InitialModel() Model {
	username := "mysterious_user"
	if sysUser, err := user.Current(); err == nil {
		username = sysUser.Username
	}

	return Model{
		SystemUsername: username,
		InputBuffer:    "",
		CursorPos:      0,
		Suggestions:    []string{},
		SelectedIndex:  -1,
		ShowDropdown:   false,
		ActiveView:     "prompt",
		SelectedProj:   "",
		Message:        "",
		History: []HistoryItem{
			{
				Command: "system --boot",
				Output:  "Be at ease",
			},
		},
		ScrollOffset:         0,
		EyeCooldownRemaining: randomEyeCooldown(),
	}
}

// ResetPrompt clears the input line and any transient UI state.
func (m Model) ResetPrompt() Model {
	m.InputBuffer = ""
	m.CursorPos = 0
	m.Suggestions = []string{}
	m.SelectedIndex = -1
	m.ShowDropdown = false
	m.ActiveView = "prompt"
	m.SelectedProj = ""
	m.Message = ""
	return m
}

func (m Model) irisTarget() IrisTarget {
	if m.FloatingEye != nil {
		return IrisTarget{X: int(m.FloatingEye.X), Y: int(m.FloatingEye.Y), Active: true}
	}
	if m.MouseX > 0 || m.MouseY > 0 {
		return IrisTarget{X: m.MouseX, Y: m.MouseY, Active: true}
	}
	return IrisTarget{}
}
