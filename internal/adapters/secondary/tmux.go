package secondary

import (
	"context"
	"os/exec"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
)

type TmuxCmds struct {
	HasSession   *exec.Cmd
	NewSession   *exec.Cmd
	SelectWindow *exec.Cmd
}

type TmuxAdapter struct {
	cfg  *entity.Config
	cmds *TmuxCmds
}

func NewTmuxAdapter(ctx context.Context, config *entity.Config) (*TmuxAdapter, error) {
	return &TmuxAdapter{
		cfg: config,
	}, nil
}

func (a *TmuxAdapter) HasSession(ctx context.Context) (int, error) {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.HasSessionCmd), "-t "+a.cfg.Session)

	if err := cmd.Run(); err != nil {
		return tmux.SessionNotExists, err
	}

	return tmux.SessionExists, nil
}

func (a *TmuxAdapter) NewSession(ctx context.Context, name string) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewSessionCmd), "-d", "-s "+a.cfg.Session, "-n "+name)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) AttachSession(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.AttachCmd), "-t "+a.cfg.Session)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) SendKeys(ctx context.Context, name string, keyCmd string) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.SendKeysCmd), "-t "+a.cfg.Session+":"+name, keyCmd, "C-m")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) NewWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewWindowCmd), "-t "+a.cfg.Session, "-n "+a.cfg.Windows[cfgIndex].Name)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (a *TmuxAdapter) SelectWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.SelectWindowCmd), "-t "+a.cfg.Session+":"+a.cfg.Windows[cfgIndex].Name)

	if err := cmd.Run(); err != nil {
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
