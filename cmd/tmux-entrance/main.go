package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"

	tmuxentrance "github.com/pigeonligh/sundries/pkg/bubbletea/models/tmux-entrance"
	"github.com/pigeonligh/sundries/pkg/utils"
)

func main() {
	model := tmuxentrance.New()
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		utils.Fatal(err)
	}

	switch model.Result.Action {
	case "shell":
		utils.SysExecShell(os.Getenv("SHELL"))
	case "tmux":
		model.Tmux.ExecSession(model.Result.Value)
	}

	os.Exit(1)
}
