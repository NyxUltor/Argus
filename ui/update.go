package ui

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"argus/autocomplete"
	"argus/data"

	tea "github.com/charmbracelet/bubbletea"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func (m Model) Init() tea.Cmd {
	return nil
}

// getCommandOutput returns the static output for built-in commands
func getCommandOutput(cmd string) string {
	store := data.GetPortfolioData()
	switch cmd {
	case "about":
		return store.AboutText
	case "skills":
		var sb strings.Builder
		sb.WriteString("--- RECOGNIZED SYSTEM CAPABILITIES ---\n")
		for _, skill := range store.Skills {
			sb.WriteString(fmt.Sprintf("* %s\n", skill))
		}
		return sb.String()
	case "projects":
		var sb strings.Builder
		sb.WriteString("--- ACTIVE BUILD PIPELINES ---\n")
		var keys []string
		for k := range store.Projects {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			proj := store.Projects[k]
			sb.WriteString(fmt.Sprintf("\n> %s [%s]\n  %s\n", proj.Name, strings.Join(proj.Stack, ", "), proj.Description))
		}
		return sb.String()
	case "contact":
		return store.ContactText
	case "help":
		return "Available system commands: about, projects, skills, contact, clear, exit, quit"
	default:
		return ""
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commandFinishedMsg:
		if msg.Index >= 0 && msg.Index < len(m.History) {
			m.History[msg.Index].Output = msg.Output
		}
		return m, nil

	case tea.KeyMsg:
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
			runes := []rune(m.InputBuffer)
			m.CursorPos = len(runes)
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
				return m, nil
			}

			// check if it's a built-in
			store := data.GetPortfolioData()
			isBuiltin := false
			builtins := append(store.Commands, "sudo", "help")
			for _, b := range builtins {
				if cmdClean == b {
					isBuiltin = true
					break
				}
			}

			var output string

			if isBuiltin {
				if cmdClean == "sudo" { // Just in case you did put it in builtins
					quotes := []string{
						"Who decided that?",
						"That level of genjutsu doesn't work on me.",
						"Since when did you fall under the illusion that you could command me?",
					}
					output = quotes[rng.Intn(len(quotes))]
					m.History = append(m.History, HistoryItem{
						Command: cmdText,
						Output:  output,
					})
					return m, nil
				}

				output = getCommandOutput(cmdClean)
				m.History = append(m.History, HistoryItem{
					Command: cmdText,
					Output:  output,
				})
				return m, nil

			} else {
				// CATCH ALL UNRECOGNIZED COMMANDS HERE

				// If they try to run a local system command with sudo, block it with style
				if cmdClean == "sudo" || strings.HasPrefix(cmdText, "sudo ") {
					quotes := []string{
						"Who decided that?",
						"That level of genjutsu doesn't work on me.",
						"Since when did you fall under the illusion that you could command me?",
					}
					output = quotes[rng.Intn(len(quotes))]
					m.History = append(m.History, HistoryItem{
						Command: cmdText,
						Output:  output,
					})
					return m, nil
				}

				// Otherwise, hand it off to your exec files (local runs it, ssh says command not found)
				return handleFallbackCommand(m, cmdText)
			}
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

	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height
	}

	return m, nil
}
