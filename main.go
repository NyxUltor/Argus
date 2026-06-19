package main

import (
	"fmt"
	"os"

	"nyx-portfolio/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create a new Bubble Tea runtime instance passing our clean model state
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
