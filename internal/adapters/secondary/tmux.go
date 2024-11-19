package secondary

import (
	"context"
	"os/exec"

	"github.com/nxdir-s/gomux/internal/core/valobj"
)

type TmuxCmds struct {
	HasSession   *exec.Cmd
	NewSession   *exec.Cmd
	SelectWindow *exec.Cmd
}

type TmuxAdapter struct {
	cfg  *valobj.Config
	cmds *TmuxCmds
}

func NewTmuxAdapter(config *valobj.Config) (*TmuxAdapter, error) {
	return &TmuxAdapter{
		cfg: config,
	}, nil
}

func (a *TmuxAdapter) SetupCmds(ctx context.Context) {
	commands := &TmuxCmds{
		HasSession:   exec.CommandContext(ctx, "tmux", "has-session", "-t "+a.cfg.Session, "2>/dev/null"),
		NewSession:   exec.CommandContext(ctx, "tmux", "new-session", "-d", "-s "+a.cfg.Session, "-n editor"),
		SelectWindow: exec.CommandContext(ctx, "tmux", "select-window", "-t "+a.cfg.Session+":editor"),
	}

	a.cmds = commands
}

func (a *TmuxAdapter) HasSession(ctx context.Context) (int, error) {
	if err := a.cmds.HasSession.Run(); err != nil {
		return 0, err
	}

	return 1, nil
}

func (a *TmuxAdapter) NewSession(ctx context.Context) error {
	if err := a.cmds.NewSession.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) AttachSession(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "tmux", "attach-session", "-t "+a.cfg.Session)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) SendKeys(ctx context.Context, windowName string, keyCmd string) error {
	cmd := exec.CommandContext(ctx, "tmux", "send-keys", "-t "+a.cfg.Session+":"+windowName, keyCmd, "C-m")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) NewWindow(ctx context.Context, name string) error {
	cmd := exec.CommandContext(ctx, "tmux", "new-window", "-t "+a.cfg.Session, "-n "+name)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) SelectWindow(ctx context.Context) error {
	if err := a.cmds.SelectWindow.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) SetOption(ctx context.Context) error {
	return nil
}

func (a *TmuxAdapter) SetWindowOpt(ctx context.Context) error {
	return nil
}
