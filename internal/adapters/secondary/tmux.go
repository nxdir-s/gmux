package secondary

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
	"github.com/nxdir-s/gomux/internal/ports"
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
	cmd ports.CommandPort
}

// NewTmuxAdapter creates a tmux adapter
func NewTmuxAdapter(config *entity.Config, cmd ports.CommandPort) (*TmuxAdapter, error) {
	return &TmuxAdapter{
		cfg: config,
		cmd: cmd,
	}, nil
}

// HasSession checks for an already existing tmux session
func (a *TmuxAdapter) HasSession(ctx context.Context) int {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.HasSessionCmd), "-t", a.cfg.Session)

	fmt.Fprintf(os.Stdout, "checking for existing session '%s'\n", a.cfg.Session)

	output, err := a.cmd.Exec(ctx, cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed: %s\n", string(tmux.HasSessionCmd), err.Error())

		return tmux.SessionNotExists
	}

	buf := &strings.Builder{}
	io.Copy(buf, output)

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.HasSessionCmd), buf.String())

	return tmux.SessionExists
}

// NewSession creates a new tmux session
func (a *TmuxAdapter) NewSession(ctx context.Context, name string) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewSessionCmd), "-d", "-s", a.cfg.Session, "-n", name)

	fmt.Fprintf(os.Stdout, "creating new session named '%s'\n", a.cfg.Session)

	output, err := a.cmd.Exec(ctx, cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed: %s\n", string(tmux.NewSessionCmd), err.Error())

		return &ErrNewSession{name, err}
	}

	buf := &strings.Builder{}
	io.Copy(buf, output)

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.NewSessionCmd), buf.String())

	return nil
}

// AttachSession attempts attaching to a tmux session
func (a *TmuxAdapter) AttachSession(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.AttachCmd), "-t", a.cfg.Session)
	cmd.Stdin = os.Stdin

	output, err := a.cmd.Exec(ctx, cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed: %s\n", string(tmux.AttachCmd), err.Error())

		return &ErrAttachSession{a.cfg.Session, err}
	}

	buf := &strings.Builder{}
	io.Copy(buf, output)

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.AttachCmd), buf.String())

	return nil
}

// SendKeys executes the config window's command
func (a *TmuxAdapter) SendKeys(ctx context.Context, cfgIndex int) error {
	cmdArgs := []string{string(tmux.SendKeysCmd), "-t", a.cfg.Session + ":" + a.cfg.Windows[cfgIndex].Name}
	cmdArgs = append(cmdArgs, a.cfg.Windows[cfgIndex].Cmd...)

	cmd := exec.CommandContext(ctx, tmux.Alias, cmdArgs...)

	output, err := a.cmd.Exec(ctx, cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed: %s\n", string(tmux.SendKeysCmd), err.Error())

		return &ErrSendKeys{strings.Join(cmdArgs, " "), err}
	}

	buf := &strings.Builder{}
	io.Copy(buf, output)

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.SendKeysCmd), buf.String())

	return nil
}

// NewWindow creates a new tmux window
func (a *TmuxAdapter) NewWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewWindowCmd), "-t", a.cfg.Session, "-n", a.cfg.Windows[cfgIndex].Name)

	output, err := a.cmd.Exec(ctx, cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed: %s\n", string(tmux.NewWindowCmd), err.Error())

		return &ErrNewWindow{a.cfg.Windows[cfgIndex].Name, err}
	}

	buf := &strings.Builder{}
	io.Copy(buf, output)

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.NewWindowCmd), buf.String())

	return nil
}

// SelectWindow selects a tmux window based on the cfgIndex
func (a *TmuxAdapter) SelectWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.SelectWindowCmd), "-t", a.cfg.Session+":"+a.cfg.Windows[cfgIndex].Name)

	output, err := a.cmd.Exec(ctx, cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s failed: %s\n", string(tmux.SelectWindowCmd), err.Error())

		return &ErrSelectWindow{a.cfg.Windows[cfgIndex].Name, err}
	}

	buf := &strings.Builder{}
	io.Copy(buf, output)

	fmt.Fprintf(os.Stdout, "%s output: %s\n", string(tmux.SelectWindowCmd), buf.String())

	return nil
}
