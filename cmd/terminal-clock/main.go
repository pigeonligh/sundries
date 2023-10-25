package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	terminalclock "github.com/pigeonligh/sundries/pkg/bubbletea/models/terminal-clock"
)

func main() {
	p := tea.NewProgram(terminalclock.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
