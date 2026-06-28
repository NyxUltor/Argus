package ui

// update.go — Bubble Tea Update loop and Init.
// Command output logic lives in commands.go.
// Message types live in messages.go.
// Layout math lives in layout.go.

import (
	"strings"
	"unicode/utf8"

	"argus/autocomplete"
	"argus/data"

	tea "github.com/charmbracelet/bubbletea"
)

func enableMouse() tea.Msg {
	return tea.EnableMouseAllMotion()
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		eyeTick(),
		enableMouse,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// --- Async command completion ---
	case commandFinishedMsg:
		if msg.Index >= 0 && msg.Index < len(m.History) {
			m.History[msg.Index].Output = msg.Output
		}
		return m, nil

	case ballTickMsg:
		if len(m.Balls) > 0 {
			m = m.updateBalls()
			return m, ballTick()
		}
		return m, nil

	case eyeTickMsg:
		m = m.updateFloatingEye()
		return m, eyeTick()

	case takeoverTickMsg:
		if m.ActiveTakeover == TakeoverMatrix {
			return m.updateMatrix()
		}
		if m.ActiveTakeover == TakeoverStarfield {
			return m.updateStarfield()
		}
		return m, nil

	// --- Mouse: scroll wheel and scrollbar click ---
	case tea.MouseMsg:
		m.MouseX = msg.X
		m.MouseY = msg.Y
		if msg.Button == tea.MouseButtonWheelUp {
			maxScroll := m.MaxScrollOffset()
			m.ScrollOffset++
			if m.ScrollOffset > maxScroll {
				m.ScrollOffset = maxScroll
			}
			return m, nil
		}
		if msg.Button == tea.MouseButtonWheelDown {
			m.ScrollOffset--
			if m.ScrollOffset < 0 {
				m.ScrollOffset = 0
			}
			return m, nil
		}
		// Scrollbar track click: right-most two columns.
		if msg.Action == tea.MouseActionRelease && msg.X >= m.TerminalWidth-2 {
			statsHeight := len(strings.Split(m.renderStatsPanel(), "\n"))
			clickY := msg.Y - statsHeight
			if clickY >= 0 {
				H_history := m.TerminalHeight - statsHeight - len(strings.Split(m.renderInputSection(), "\n")) - 2
				if H_history > 0 {
					ratio := float64(clickY) / float64(H_history)
					if ratio < 0 {
						ratio = 0
					}
					if ratio > 1 {
						ratio = 1
					}
					max := m.MaxScrollOffset()
					m.ScrollOffset = max - int(ratio*float64(max)+0.5)
					if m.ScrollOffset < 0 {
						m.ScrollOffset = 0
					}
					if m.ScrollOffset > max {
						m.ScrollOffset = max
					}
				}
			}
		}
		return m, nil

	// --- Keyboard ---
	case tea.KeyMsg:
		// Takeover: any key exits back to portfolio. Key is consumed.
		if m.ActiveTakeover != TakeoverNone {
			m = m.clearTakeover()
			return m, nil
		}
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "left":
			if m.CursorPos > 0 {
				m.CursorPos--
			}
			return m, nil

		case "right":
			runes := []rune(m.InputBuffer)
			if m.CursorPos < len(runes) {
				m.CursorPos++
			}
			return m, nil

		case "home", "ctrl+a":
			m.CursorPos = 0
			return m, nil

		case "end", "ctrl+e":
			m.CursorPos = len([]rune(m.InputBuffer))
			return m, nil

		case "backspace":
			runes := []rune(m.InputBuffer)
			if m.CursorPos > 0 {
				m.InputBuffer = string(append(runes[:m.CursorPos-1], runes[m.CursorPos:]...))
				m.CursorPos--
				m.Suggestions = autocomplete.FilterCommands(m.InputBuffer)
			}
			return m, nil

		case "delete":
			runes := []rune(m.InputBuffer)
			if m.CursorPos < len(runes) {
				m.InputBuffer = string(append(runes[:m.CursorPos], runes[m.CursorPos+1:]...))
				m.Suggestions = autocomplete.FilterCommands(m.InputBuffer)
			}
			return m, nil

		case "enter":
			return m.handleEnter()

		default:
			if utf8.RuneCountInString(msg.String()) == 1 {
				runes := []rune(m.InputBuffer)
				newRune := []rune(msg.String())[0]
				if m.CursorPos >= len(runes) {
					runes = append(runes, newRune)
				} else {
					runes = append(runes[:m.CursorPos], append([]rune{newRune}, runes[m.CursorPos:]...)...)
				}
				m.InputBuffer = string(runes)
				m.CursorPos++
				m.Suggestions = autocomplete.FilterCommands(m.InputBuffer)
			}
		}

	// --- Terminal resize ---
	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height
		// Bug 3 fix: clamp ScrollOffset to the newly recalculated boundary.
		maxScroll := m.MaxScrollOffset()
		if m.ScrollOffset > maxScroll {
			m.ScrollOffset = maxScroll
		}
		if m.ScrollOffset < 0 {
			m.ScrollOffset = 0
		}
	}

	return m, nil
}

// handleEnter processes a submitted command, dispatching to built-ins or fallback exec.
func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	m.ScrollOffset = 0
	cmdText := m.InputBuffer
	cmdClean := strings.TrimSpace(strings.ToLower(cmdText))
	m.InputBuffer = ""
	m.CursorPos = 0
	m.Suggestions = []string{}

	if cmdClean == "" {
		return m, nil
	}
	if cmdClean == "exit" || cmdClean == "quit" {
		return m, tea.Quit
	}
	if cmdClean == "clear" {
		m.History = []HistoryItem{}
		m.Balls = []Ball{}
		return m, nil
	}

	// summon dispatch
	if strings.HasPrefix(cmdClean, "summon ") {
		return m.handleSummon(cmdClean)
	}

	// 'over' terminates active overlay
	if cmdClean == "over" {
		if len(m.Balls) == 0 && m.FloatingEye == nil {
			m.History = append(m.History, HistoryItem{
				Command: cmdText,
				Output:  "No active overlay to dismiss.",
			})
			return m, nil
		}
		m.Balls = []Ball{}
		return m, nil
	}

	// Any command during overlay clears balls before executing
	if len(m.Balls) > 0 {
		m.Balls = []Ball{}
	}
	// Any command during takeover clears it
	if m.ActiveTakeover != TakeoverNone {
		m = m.clearTakeover()
	}

	// Sudo intercept (applies everywhere, before built-in check).
	if isSudoAttempt(cmdClean, cmdText) {
		m.History = append(m.History, HistoryItem{
			Command: cmdText,
			Output:  randomSudoResponse(),
		})
		return m, nil
	}

	// Built-in command lookup.
	store := data.GetPortfolioData()
	for _, b := range store.Commands {
		if cmdClean == b {
			output := getCommandOutput(cmdClean)
			m.History = append(m.History, HistoryItem{
				Command: cmdText,
				Output:  output,
			})
			return m, nil
		}
	}

	// Unrecognised — fall through to build-tag-selected exec handler.
	return handleFallbackCommand(m, cmdText)
}
