package tmuxentrance

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pigeonligh/sundries/pkg/bubbletea/render"
	"github.com/pigeonligh/sundries/pkg/tmux"
)

type Result struct {
	Action string
	Value  string
}

type model struct {
	width  int
	height int

	choices []tmux.TmuxSession
	current int

	Result Result

	Tmux *tmux.TmuxHelper
}

func New() *model {
	return &model{
		choices: make([]tmux.TmuxSession, 0),
		Tmux:    tmux.New(os.Getenv("TMUX_SESSION")),
	}
}

func (m *model) Init() tea.Cmd {
	m.choices = m.Tmux.ListSessions(true)
	m.current = len(m.choices) - 1

	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.Result = Result{
				Action: "shell",
			}
			return m, tea.Quit
		case "ctrl+c":
			return m, tea.Quit
		case "up", "j":
			if m.current > 0 {
				m.current--
			}
			return m, nil
		case "down", "k":
			if m.current < len(m.choices)-1 {
				m.current++
			}
			return m, nil
		case "enter":
			m.Result = Result{
				Action: "tmux",
				Value:  m.choices[m.current].Name,
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m *model) View() string {
	s0 := fmt.Sprintf("\nWelcome, %s!\n", os.Getenv("USER"))
	s1 := "Choose your session:\n\n"

	for i, choice := range m.choices {
		if i == m.current {
			s1 = s1 + "> " + choice.Title + "\n"
		} else {
			s1 = s1 + "  " + choice.Title + "\n"
		}
	}

	if renderred, ok := render.CenterRender.Render(m.width, m.height, s0, s1); ok {
		return renderred
	}
	return s0
}
