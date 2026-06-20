package ui

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

	History []HistoryItem
}

func InitialModel() Model {
	return Model{
		InputBuffer:   "",
		CursorPos:     0,
		Suggestions:   []string{},
		SelectedIndex: -1,
		ShowDropdown:  false,
		ActiveView:    "prompt",
		SelectedProj:  "",
		Message:       "",
		History: []HistoryItem{
			{
				Command: "system --boot",
				Output:  "Argus Terminal Controller initialized.\nType 'help' to see available commands.",
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
