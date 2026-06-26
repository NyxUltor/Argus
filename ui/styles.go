package ui

// styles.go — centralised Lipgloss colour palette and style definitions.
// To retheme Argus, only this file needs to be edited.

import "github.com/charmbracelet/lipgloss"

var (
	// Header / logo
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ffd7")) // vibrant cyan

	// Stats panel labels and values
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff79c6")).Bold(true) // Dracula pink
	valStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#f8f8f2"))            // Dracula white

	// Structural borders / dividers
	borderStyle1 = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9")) // Dracula purple (scrollbar handle)
	borderStyle2 = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4")) // Dracula grey  (scrollbar track / divider)

	// Input prompt
	promptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true) // Dracula green

	// Cursor block (reverse-video)
	cursorStyle = lipgloss.NewStyle().Reverse(true)

	// Accent styles used inline
	accentGreen = lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true)
)
