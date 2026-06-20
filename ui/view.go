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

// wrapText wraps the given text based on character length.
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
		runes := []rune(line)
		for len(runes) > width {
			result = append(result, string(runes[:width]))
			runes = runes[width:]
		}
		if len(runes) > 0 {
			result = append(result, string(runes))
		}
	}
	return result
}

func (m Model) renderStatsPanel() string {
	logoLines := []string{
		"    ▄▄▄▄▄▄▄    ",
		"  ▄▀       ▀▄  ",
		" █   ▄▄▄▄▄   █ ",
		" █  █  █  █  █ ",
		" █   ▀▀▀▀▀   █ ",
		"  ▀▄       ▄▀  ",
		"    ▀▀▀▀▀▀▀    ",
	}

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
		labelStyle.Render("Uptime:     ") + valStyle.Render(uptimeStr),
		labelStyle.Render("Terminal:   ") + valStyle.Render(fmt.Sprintf("%dx%d", m.TerminalWidth, m.TerminalHeight)),
		labelStyle.Render("Shell/Host: ") + valStyle.Render(hostMode),
		labelStyle.Render("Commands:   ") + valStyle.Render("about, projects, skills, contact, clear, exit, quit"),
		labelStyle.Render("Engine:     ") + valStyle.Render("Go / Bubble Tea / Lip Gloss"),
		labelStyle.Render("Status:     ") + lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true).Render("ONLINE"),
	}

	var lines []string
	for i := 0; i < len(logoLines); i++ {
		logoPart := logoStyle.Render(logoLines[i])
		statPart := stats[i]
		lines = append(lines, fmt.Sprintf("  %s   %s", logoPart, statPart))
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

func (m Model) renderInputLine() string {
	prompt := promptStyle.Render("[veryl@ARGUS ~]$ ")
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
		// Render prompt line
		promptLine := promptStyle.Render("[veryl@ARGUS ~]$ ") + item.Command
		lines = append(lines, promptLine)

		// Render output lines
		if item.Output != "" {
			outputLines := wrapText(item.Output, m.TerminalWidth)
			lines = append(lines, outputLines...)
		}
	}
	return lines
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
	// Separators:
	// - Border 1 (under stats panel): 1 line
	// - Border 2 (above input line): 1 line
	H_history := m.TerminalHeight - H_top - H_input - 2
	if H_history < 0 {
		H_history = 0
	}

	historyLines := m.renderHistoryLines()

	var historySection string
	if len(historyLines) > H_history {
		historySection = strings.Join(historyLines[len(historyLines)-H_history:], "\n")
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
	s.WriteString(borderStyle1.Render(strings.Repeat("═", m.TerminalWidth)))
	s.WriteString("\n")
	s.WriteString(historySection)
	s.WriteString("\n")
	s.WriteString(borderStyle2.Render(strings.Repeat("─", m.TerminalWidth)))
	s.WriteString("\n")
	s.WriteString(inputSection)

	return s.String()
}
