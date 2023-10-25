package terminalclock

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pigeonligh/sundries/pkg/bubbletea/render"
)

type tickMsg struct{}

type model struct {
	width  int
	height int
}

func New() *model {
	return &model{}
}

func (m *model) Init() tea.Cmd {
	return m.tickCmd()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tickMsg:
		return m, m.tickCmd()
	}
	return m, nil
}

func (m model) View() string {
	s := []string{fmt.Sprintf("Welcome, %s!", os.Getenv("USER"))}
	s = append(s, "Time: "+time.Now().Format("2006-01-02 15:04:05"))
	if renderred, ok := render.CenterRender.Render(m.width, m.height, s...); ok {
		return renderred
	}
	return s[0]
}

func (m model) tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
