package ui

import "os/user"

type HistoryItem struct {
	Command string
	Output  string
}

type commandFinishedMsg struct {
	Index  int
	Output string
}

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

	History []HistoryItem
}

func InitialModel() Model {
	// 1. Create a default fallback name
	username := "mysterious_user"

	// 2. Try to get the real OS username
	if sysUser, err := user.Current(); err == nil {
		username = sysUser.Username
	}

	// 3. Return your model using that username variable
	return Model{
		SystemUsername: username, // <-- Works perfectly every time now!
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
	}
}

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
