package ui

// render_header.go — renders the top stats/dashboard panel.
// Logo art and owner metadata come from data/portfolio.go (user-editable).

import (
	"fmt"
	"strings"
	"time"

	"argus/data"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderStatsPanel() string {
	store := data.GetPortfolioData()

	uptime := time.Since(store.ProjectStart)
	days := int(uptime.Hours()) / 24
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60
	uptimeStr := fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)

	hostMode := getHostModeString()

	stats := []string{
		accentGreen.Render("ARGUS CORE CONTROLLER v1.0.0"),
		labelStyle.Render("Project age:     ") + valStyle.Render(uptimeStr),
		labelStyle.Render("Terminal size:   ") + valStyle.Render(fmt.Sprintf("%dx%d", m.TerminalWidth, m.TerminalHeight)),
		labelStyle.Render("Shell/Host:      ") + valStyle.Render(hostMode),
		labelStyle.Render("github:          ") + valStyle.Render(store.OwnerHandle),
		labelStyle.Render("Status:          ") + lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true).Render("ONLINE"),
	}

	var lines []string
	art := applyIrisOffset(EyeArt, m.irisTarget(), m.TerminalWidth, m.TerminalHeight)
	
	maxLogoWidth := 0
	for _, line := range art {
		l := len([]rune(line))
		if l > maxLogoWidth {
			maxLogoWidth = l
		}
	}

	for i := 0; i < len(art); i++ {
		lineRune := []rune(art[i])
		padLogo := maxLogoWidth - len(lineRune)
		paddedLogoLine := string(lineRune)
		if padLogo > 0 {
			paddedLogoLine += strings.Repeat(" ", padLogo)
		}
		logoPart := logoStyle.Render(paddedLogoLine)

		if i < len(stats) {
			visWidthStats := lipgloss.Width(stats[i])
			spacerWidth := m.TerminalWidth - 2 - maxLogoWidth - visWidthStats - 4
			if spacerWidth < 3 {
				spacerWidth = 3
			}
			lines = append(lines, fmt.Sprintf("  %s%s%s", logoPart, strings.Repeat(" ", spacerWidth), stats[i]))
		} else {
			lines = append(lines, fmt.Sprintf("  %s", logoPart))
		}
	}
	return strings.Join(lines, "\n")
}
