package ui

type Model struct {
	InputBuffer    string   
	CursorPos      int      
	Suggestions    []string 
	SelectedIndex  int      
	ShowDropdown   bool     
	ActiveView     string   
	SelectedProj   string   
	TerminalWidth  int
	TerminalHeight int
}

func InitialModel() Model {
	return Model{
		InputBuffer:   "",
		CursorPos:     0,
		Suggestions:   []string{},
		SelectedIndex: -1,
		ShowDropdown:  false,
		ActiveView:    "prompt",
	}
}

func (m *Model) ResetPrompt() {
	m.InputBuffer = ""
	m.CursorPos = 0
	m.Suggestions = []string{}
	m.SelectedIndex = -1
	m.ShowDropdown = false
	m.ActiveView = "prompt"
	m.SelectedProj = ""
}