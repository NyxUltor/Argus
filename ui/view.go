package ui

import (
	"fmt"
	"strings"
	"time"

	"argus/data"

	"github.com/charmbracelet/lipgloss"
)

var (
	logoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ffd7"))            // vibrant cyan
	labelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff79c6")).Bold(true) // Dracula pink
	valStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#f8f8f2"))            // Dracula white
	borderStyle1 = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9"))            // Dracula purple
	borderStyle2 = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4"))            // Dracula grey
	promptStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true) // Dracula green
	cursorStyle  = lipgloss.NewStyle().Reverse(true)
)

func wrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		if len(line) == 0 {
			result = append(result, "")
			continue
		}
		words := strings.Fields(line)
		current := ""
		for _, word := range words {
			wordRunes := []rune(word)
			spaceNeeded := 0
			if current != "" {
				spaceNeeded = 1
			}
			if len([]rune(current))+spaceNeeded+len(wordRunes) > width {
				if current != "" {
					result = append(result, current)
					current = ""
				}
				for len(wordRunes) > width {
					result = append(result, string(wordRunes[:width]))
					wordRunes = wordRunes[width:]
				}
				current = string(wordRunes)
			} else {
				if current == "" {
					current = word
				} else {
					current += " " + word
				}
			}
		}
		if current != "" {
			result = append(result, current)
		}
	}
	return result
}

func (m Model) renderStatsPanel() string {
	logoLines := []string{
		"⠀⠀⠀⠀ ⢀⠀⠀⠀⠀⠀⠀⠀⠄⠀⠀⠀⠤⢀⣁⣀⠠⣤⠌⢻⣛⣷⣖⣿⣿⣿⣗⣷⠖⣿⣿⣿⣿⣻⣿⠯⣡⠔⢁⠔⣠⠞⣡⣔⠼⣶⠟⠮⣃⠐⠉⠁⠀⠀⠀⠉⠁⠀⠀⠠⡠⠤⠲⠊⡀⠄⠁⠀⠀  ⡄⣲",
		"⠀⠀⠀ ⡠⠁⠀⠀⠀⠀⠀⠈⠑⠒⠈⢩⣶⣶⣶⣾⢲⣛⣿⢿⠟⣿⣿⣿⣿⡥⣶⡷⣟⣿⣿⣿⣿⡥⣚⣥⡾⣛⣛⣴⡿⠎⣫⠟⠉⠐⠂⢸⠐⠀⠀⠀⠀⠀⠀⠉⠁⠒⠑⠒⠀⠈⠀⠀  ⠀⢀⣴⣬⣾⣿⣋ ",
		"⠀⢀⠔⠁⠀⠀⠀⢀⠀⠀⢀⣐⣛⢻⣿⠿⠿⠿⣽⡷⣯⣿⣶⣞⣭⣯⠷⣻⣽⣷⣿⣿⣿⡿⢛⠁⢅⣋⣝⣶⠶⠏⠡⠄⠒⠀⢈⠄⠀⠀⡠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀  ⠀⠀⠀⣀⢐⣴⣯⣿⣿⠟⣩⡿   ",
		"⠀⠀⡤⠀⠀⢪⡠⠤⠬⢭⣭⢿⣟⣷⣽⣛⣛⣭⡿⣟⡿⢿⢿⣿⡿⣾⠻⠍⢓⠩⠕⠏⠉⠋⠐⠀⠂⠀⠀⠀⠀⠠⠈⠀⠀⣴⠀⠀⢀⠀⠀⣀⢀⣀⣠⣠⣤⣤⣴⣶⣶⣶⣶⣿⣿⣿⣿⣿⣿⣿⣟⣠⣾⠓⠀   ",
		"⠀⣁⡤⠟⠙⢁⢔⠾⢫⢴⣯⣟⣿⠞⡝⠉⠉⠋⠃⠊⠁⠀⠀⣀⣅⣀⣀⣀⣀⣀⢠⣶⣠⣤⢤⣴⠶⣶⠾⠥⣴⡶⣶⣾⣭⣟⠺⡿⢟⠻⠟⠛⠻⣿⢿⣿⣟⣛⣭⣴⣾⣾⣿⣿⠿⠟⠋⠉⠁⠀       ⠀ ",
		"⠜⡁⠤⢒⢀⣄⣴⠺⠛⠛⠑⣉⢠⡴⣶⣒⣢⣦⣦⣴⣾⣷⣿⣯⣿⣿⣿⣿⣿⣿⣼⣥⣤⣾⣿⣶⣾⣿⣶⣾⣴⣿⣷⣾⣶⣿⣿⣾⣷⣿⣿⣿⣿⣿⣿⣿⣿⡋⠉⢀⡠⠆⣀⣀⣀⣀⣀⠠           ",
		"⡰⢖⣿⣻⠏⠛⠁⠀⠀⠠⢠⣶⣭⣾⢿⣛⣿⡯⣴⣷⣷⣿⣿⣿⣿⣿⣿⣿⣿⡟⢿⢿⣿⣿⣿⣿⣾⣿⣿⣿⡿⢿⣿⢾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣮⣙⣾⣿⡋⠉⠃⠉⠀⠀⠀⠀⠀           ",
		"⠋⠉⠀⠀⠀⠀⠀⢄⣠⣴⠿⢫⣿⡷⣿⡹⠑⣱⣿⣿⣿⣋⠝⢋⡛⢿⣿⣿⣽⣿⣿⣾⡿⠻⣿⡟⢩⢻⣉⣿⢋⢂⢀⠄⢈⠋⢟⠟⢿⣿⣿⣿⢿⣿⢿⡿⣿⠫⣈⠑⠊⠉⠉⠉⠉⠉⠉⠉⠉          ",
		"⡀⠀⠀⠀⠀⠀⣤⣾⢿⠿⣤⣾⡿⣱⣟⣷⣶⣿⡻⠹⣿⣏⠀⠨⠷⣼⣿⠿⣿⠿⠿⠿⠓⡖⢯⣣⡈⠨⣽⠏⠈⠈⠊⠀⠈⠂⠘⠆⣬⡿⠋⢩⡘⠟⠚⠓⠿⡦⣉⠂⡀⠀⠀⠀⠀⠀⠀⠀			",
		"⠀⠀⠀⢀⡴⡾⡿⠃⣠⣼⣿⢿⡟⠛⠓⠒⠀⠀⠀⠀⠈⢷⣦⡄⠁⠈⠁⢀⠁⠀⠀⠐⠀⠐⢈⢊⢮⣿⠏⠀⠀⠀⠀⠀⠀⠀⢰⡽⠟⠀⣶⣿⡿⣇⠉⠛⢬⣧⠂⠑⠂⠁⠀⠀⠀⠀⠀⠀			",
		"⠀⠀⢀⢽⡾⠋⢠⣾⢟⡾⠟⡟⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⢿⣷⣤⡤⢀⢀⠠⢀⣀⡀⢀⣠⠶⠋⠀⠀⠀⠀⠀⠀⣀⣤⠖⠋⠀⣠⣾⣯⣟⣛⣾⡆⠀⠀⠹⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀			",
		"⠀⠀⡡⢉⣤⢞⡽⠒⠋⣴⠬⠼⠓⠒⠦⠤⠤⣀⣀⡀⠀⠀⠀⠀⠀⠛⠛⠛⠗⠿⠿⠛⠛⠋⠁⠀⠀⠀⠀⡀⠄⠐⠈⠀⢀⣠⣶⣿⣿⣷⡃⠈⢉⢼⠗⠀⠀⠰⠐⠀⠀⠀⠀⠀⠀⠀⠀⠀			",
		"⠀⠀⢔⡿⢵⠟⠒⢉⣡⡴⠵⠒⠚⣲⠤⢖⡀⡀⠤⠈⠚⠛⠚⠶⠤⠤⠤⠄⠀⠀⠄⢀⠀⠀⠔⠒⠂⠉⠀⢀⣀⣤⣰⣵⣿⢿⣿⠿⣿⣿⣯⢊⠼⡋⣹⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀			",
		"⠀⠀⠡⠀⠈⠀⠐⠁⠀⠀⠀⠐⠾⣭⣶⠉⣼⡽⠛⡥⣢⢤⣤⣀⡀⠀⠀⠀⠀⠀⠀⡀⣀⢀⢀⣀⢠⣖⣽⡼⡧⣙⠿⢿⣷⡧⢉⣾⢠⡺⣿⠏⢈⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀			",
		"⠀⠀⠁⡘⠀⠀⢄⠀⠠⢀⠀⢂⢼⠏⠀⣹⡟⠠⠀⣰⣘⣹⣡⡞⠛⣦⢱⢆⢑⢶⢤⡈⡖⣗⡷⣷⣥⡇⠈⣧⡇⢀⣠⣦⢿⣟⠋⠠⠳⠍⣼⠁⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀			",
		"⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⡠⢚⠋⠀⢰⠋⡆⠀⠀⠓⠁⣿⡏⣀⠀⠘⣞⡾⠃⢸⢿⢾⣡⣿⠄⠐⢽⡗⡀⣻⣇⡧⠟⢛⢉⡗⠀⠀⠀⢀⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀			",
	}
	// idea for later on, mirrored image on left side,constraint in mind,terminal resizing

	store := data.GetPortfolioData()
	uptime := time.Since(store.ProjectStart)
	days := int(uptime.Hours()) / 24
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60

	uptimeStr := fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)

	hostMode := getHostModeString()

	stats := []string{
		lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true).Render("ARGUS CORE CONTROLLER v1.0.0"),
		labelStyle.Render("Project age:     ") + valStyle.Render(uptimeStr),
		labelStyle.Render("Terminal size:   ") + valStyle.Render(fmt.Sprintf("%dx%d", m.TerminalWidth, m.TerminalHeight)),
		labelStyle.Render("Shell/Host: ") + valStyle.Render(hostMode),

		labelStyle.Render("github:     ") + valStyle.Render("NyxUltor"),
		labelStyle.Render("Status:     ") + lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true).Render("ONLINE"),
	}

	var lines []string
	for i := 0; i < len(logoLines); i++ {
		logoPart := logoStyle.Render(logoLines[i])
		if i < len(stats) {
			statPart := stats[i]
			lines = append(lines, fmt.Sprintf("  %s   %s", logoPart, statPart))
		} else {
			lines = append(lines, fmt.Sprintf("  %s", logoPart))
		}
	}
	return strings.Join(lines, "\n")
}

func (m Model) renderSuggestions() string {
	if len(m.Suggestions) == 0 {
		return ""
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4")).Italic(true).Render(
		fmt.Sprintf("  (Suggestions: %s)", strings.Join(m.Suggestions, ", ")),
	)
}

// currently the suggestions are only displayed they should also allow for selection and auto-completion of the input buffer when selected through navigration keys.

func (m Model) renderInputLine() string {
	// Dynamically uses m.SystemUsername for the input prompt
	prompt := promptStyle.Render(fmt.Sprintf("[%s@ARGUS ~]$ ", m.SystemUsername))
	runes := []rune(m.InputBuffer)

	cursorPos := m.CursorPos
	if cursorPos < 0 {
		cursorPos = 0
	}
	if cursorPos > len(runes) {
		cursorPos = len(runes)
	}

	var styledInput string
	if cursorPos == len(runes) {
		styledInput = string(runes) + cursorStyle.Render(" ")
	} else {
		before := string(runes[:cursorPos])
		charAtCursor := string(runes[cursorPos])
		after := string(runes[cursorPos+1:])
		styledInput = before + cursorStyle.Render(charAtCursor) + after
	}

	return prompt + styledInput
}

func (m Model) renderInputSection() string {
	var lines []string
	sug := m.renderSuggestions()
	if sug != "" {
		lines = append(lines, sug)
	}
	lines = append(lines, m.renderInputLine())
	return strings.Join(lines, "\n")
}

func (m Model) renderHistoryLines() []string {
	var lines []string
	for _, item := range m.History {
		// Dynamically uses m.SystemUsername for historical command prompts too!
		promptLine := promptStyle.Render(fmt.Sprintf("[%s@ARGUS ~]$ ", m.SystemUsername)) + item.Command
		lines = append(lines, promptLine)

		// Render output lines
		if item.Output != "" {
			outputLines := wrapText(item.Output, m.TerminalWidth-2)
			lines = append(lines, outputLines...)
		}
	}
	return lines
}

func (m Model) MaxScrollOffset() int {
	H_history := m.TerminalHeight - len(strings.Split(m.renderStatsPanel(), "\n")) - len(strings.Split(m.renderInputSection(), "\n")) - 2
	maxScroll := len(m.renderHistoryLines()) - H_history
	if maxScroll < 0 {
		return 0
	}
	return maxScroll
}

func (m Model) View() string {
	if m.TerminalWidth == 0 || m.TerminalHeight == 0 {
		return "Initializing screen buffer..."
	}

	// 1. Build the top stats panel
	topSection := m.renderStatsPanel()
	topLines := strings.Split(topSection, "\n")
	H_top := len(topLines)

	// 2. Build the bottom input section
	inputSection := m.renderInputSection()
	inputLines := strings.Split(inputSection, "\n")
	H_input := len(inputLines)

	// 3. Spacing and height calculations for scrolling history
	H_history := m.TerminalHeight - H_top - H_input - 2
	if H_history < 0 {
		H_history = 0
	}

	historyLines := m.renderHistoryLines()

	var historySection string
	if len(historyLines) > H_history {
		start := len(historyLines) - H_history - m.ScrollOffset
		end := len(historyLines) - m.ScrollOffset
		if start < 0 {
			start = 0
		}
		if end > len(historyLines) {
			end = len(historyLines)
		}
		visibleLines := historyLines[start:end]

		maxScroll := len(historyLines) - H_history
		ratio := 0.0
		if maxScroll > 0 {
			ratio = float64(m.ScrollOffset) / float64(maxScroll)
		}
		if ratio < 0 {
			ratio = 0
		}
		if ratio > 1 {
			ratio = 1
		}

		handlePos := 0
		if H_history > 1 {
			handlePos = (H_history - 1) - int(ratio*float64(H_history-1)+0.5)
		}

		finalLines := make([]string, H_history)
		for i := 0; i < H_history; i++ {
			if i < len(visibleLines) {
				finalLines[i] = visibleLines[i]
			} else {
				padLen := m.TerminalWidth - 2
				finalLines[i] = strings.Repeat(" ", padLen)
			}
			visWidth := lipgloss.Width(finalLines[i])
			padLen := (m.TerminalWidth - 2) - visWidth
			if padLen > 0 {
				finalLines[i] += strings.Repeat(" ", padLen)
			}
			if i == handlePos {
				finalLines[i] += " " + borderStyle1.Render("█")
			} else {
				finalLines[i] += " " + borderStyle2.Render("│")
			}
		}
		historySection = strings.Join(finalLines, "\n")
	} else {
		// Pad with blank lines at the top of history to keep content stacked towards bottom
		paddingCount := H_history - len(historyLines)
		padding := make([]string, paddingCount)
		for i := range padding {
			padding[i] = ""
		}
		allLines := append(padding, historyLines...)
		historySection = strings.Join(allLines, "\n")
	}

	var s strings.Builder
	s.WriteString(topSection)
	s.WriteString("\n")
	// s.WriteString(borderStyle1.Render(strings.Repeat("═", m.TerminalWidth)))
	// s.WriteString("\n")
	s.WriteString(historySection)
	s.WriteString("\n")
	s.WriteString(borderStyle2.Render(strings.Repeat("─", m.TerminalWidth)))
	s.WriteString("\n")
	s.WriteString(inputSection)

	return s.String()
}
