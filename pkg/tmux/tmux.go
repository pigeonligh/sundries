package tmux

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/pigeonligh/sundries/pkg/utils"
)

type TmuxSession struct {
	Name  string
	Title string
}

type TmuxHelper struct {
	executable string
	shell      string
	socket     string
}

func New(socket string) *TmuxHelper {
	executable, err := exec.LookPath("tmux")
	if err != nil {
		utils.Fatal(err)
	}

	shell := os.Getenv("SHELL")

	if socket == "" {
		socket = "default"
	}

	return &TmuxHelper{
		executable: executable,
		shell:      shell,
		socket:     socket,
	}
}

func (t *TmuxHelper) Command(args ...string) string {
	return strings.Join(append([]string{
		t.executable, "-L", t.socket,
	}, args...), " ")
}

func (t *TmuxHelper) ListSessions(withNew bool) []TmuxSession {
	out, _, _ := utils.ExecCommand(context.Background(), t.Command("ls"))

	sessions := make([]TmuxSession, 0)

	for _, line := range strings.Split(out, "\n") {
		name := getName(line)
		if name == "" {
			continue
		}
		sessions = append(sessions, TmuxSession{
			Name:  name,
			Title: line,
		})
	}

	if withNew {
		sessions = append(sessions, TmuxSession{
			Name:  "",
			Title: "new session",
		})
	}

	return sessions
}

func (t *TmuxHelper) ExecSession(session string) {
	env := append(os.Environ(), "TERM=xterm-256color")
	var cmd []string
	if session == "" {
		cmd = append(strings.Split(t.Command(), " "), "new", t.shell)
	} else {
		cmd = append(strings.Split(t.Command(), " "), "attach", "-t", session)
	}
	err := syscall.Exec(t.executable, cmd, env)
	if err != nil {
		panic(err)
	}
}

func getName(raw string) string {
	fields := strings.SplitN(raw, ":", 2)
	if len(fields) == 2 {
		return fields[0]
	}
	return ""
}
