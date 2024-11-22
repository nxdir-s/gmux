package secondary

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
)

type ErrNewSession struct {
	session string
	err     error
}

func (e *ErrNewSession) Error() string {
	return "error creating new session named '" + e.session + "': " + e.err.Error()
}

type ErrAttachSession struct {
	session string
	err     error
}

func (e *ErrAttachSession) Error() string {
	return "error attaching to session '" + e.session + "': " + e.err.Error()
}

type ErrNewWindow struct {
	window string
	err    error
}

func (e *ErrNewWindow) Error() string {
	return "error creating new window named '" + e.window + "': " + e.err.Error()
}

type ErrSelectWindow struct {
	window string
	err    error
}

func (e *ErrSelectWindow) Error() string {
	return "error selecting " + e.window + " window: " + e.err.Error()
}

type ErrSendKeys struct {
	cmd string
	err error
}

func (e *ErrSendKeys) Error() string {
	return "error executing " + string(tmux.SendKeysCmd) + " with cmd '" + e.cmd + "': " + e.err.Error()
}

type TmuxAdapter struct {
	cfg *entity.Config
}

func NewTmuxAdapter(config *entity.Config) (*TmuxAdapter, error) {
	return &TmuxAdapter{
		cfg: config,
	}, nil
}

func (a *TmuxAdapter) HasSession(ctx context.Context) int {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.HasSessionCmd), "-t", a.cfg.Session)

	fmt.Fprintf(os.Stdout, "checking for existing session '%s'\n", a.cfg.Session)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.HasSessionCmd), string(output))

		return tmux.SessionNotExists
	}

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.HasSessionCmd), string(output))

	return tmux.SessionExists
}

func (a *TmuxAdapter) NewSession(ctx context.Context, name string) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewSessionCmd), "-d", "-s", a.cfg.Session, "-n", name)

	fmt.Fprintf(os.Stdout, "creating new session named '%s'\n", a.cfg.Session)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed, output: %s\n", string(tmux.NewSessionCmd), string(output))

		return &ErrNewSession{name, err}
	}

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.NewSessionCmd), string(output))

	return nil
}

func (a *TmuxAdapter) AttachSession(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.AttachCmd), "-t", a.cfg.Session)
	cmd.Stdin = os.Stdin

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed, output: %s\n", string(tmux.AttachCmd), string(output))

		return &ErrAttachSession{a.cfg.Session, err}
	}

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.AttachCmd), string(output))

	return nil
}

func (a *TmuxAdapter) SendKeys(ctx context.Context, name string, args ...string) error {
	cmdArgs := []string{string(tmux.SendKeysCmd), "-t", a.cfg.Session + ":" + name}
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.CommandContext(ctx, tmux.Alias, cmdArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed, output: %s\n", string(tmux.SendKeysCmd), string(output))

		return &ErrSendKeys{strings.Join(cmdArgs, " "), err}
	}

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.SendKeysCmd), string(output))

	return nil
}

func (a *TmuxAdapter) NewWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewWindowCmd), "-t", a.cfg.Session, "-n", a.cfg.Windows[cfgIndex].Name)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed, output: %s\n", string(tmux.NewWindowCmd), string(output))

		return &ErrNewWindow{a.cfg.Windows[cfgIndex].Name, err}
	}

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.NewWindowCmd), string(output))

	return nil
}

func (a *TmuxAdapter) SelectWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.SelectWindowCmd), "-t", a.cfg.Session+":"+a.cfg.Windows[cfgIndex].Name)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed, output: %s\n", string(tmux.SelectWindowCmd), string(output))

		return &ErrSelectWindow{a.cfg.Windows[cfgIndex].Name, err}
	}

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.SelectWindowCmd), string(output))

	return nil
}
