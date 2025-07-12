package domain

import (
	"context"

	"github.com/nxdir-s/gmux/internal/core/valobj"
	"github.com/nxdir-s/gmux/internal/ports"
)

const FileName string = ".gmux.toml"

const (
	TmuxSessionExists int = iota
	TmuxSessionNotExists
)

const (
	TmuxEnterCmd        string = "C-m"
	TmuxHasSessionCmd   string = "has-session"
	TmuxNewSessionCmd   string = "new-session"
	TmuxNewWindowCmd    string = "new-window"
	TmuxSelectWindowCmd string = "select-window"
	TmuxAttachCmd       string = "attach-session"
	TmuxSendKeysCmd     string = "send-keys"
)

type ErrSessionSetup struct {
	err error
}

func (e *ErrSessionSetup) Error() string {
	return "failed to setup session: " + e.err.Error()
}

type ErrWindowSetup struct {
	err error
}

func (e *ErrWindowSetup) Error() string {
	return "failed to setup window: " + e.err.Error()
}

type Terminal struct {
	tmux ports.Tmux
	cfg  valobj.Config
}

// NewTerminal creates a Terminal orchestrator
func NewTerminal(config *valobj.Config, adapter ports.Tmux) *Terminal {
	return &Terminal{
		tmux: adapter,
		cfg:  *config,
	}
}

// StartTmux orchestrates tmux setup
func (d *Terminal) StartTmux(ctx context.Context) error {
	if exists := d.tmux.SessionExists(ctx); exists == TmuxSessionExists {
		if err := d.SetupSession(ctx); err != nil {
			return err
		}
	}

	if err := d.tmux.AttachSession(ctx); err != nil {
		return err
	}

	return nil
}

// SetupSession creates a new tmux session and windows using the config
func (d *Terminal) SetupSession(ctx context.Context) error {
	if err := d.tmux.NewSession(ctx, d.cfg.Windows[d.cfg.StartIndex].Name); err != nil {
		return &ErrSessionSetup{err}
	}

	for index := range d.cfg.Windows {
		if err := d.SetupWindow(ctx, index); err != nil {
			return &ErrSessionSetup{err}
		}
	}

	if err := d.tmux.SelectWindow(ctx, d.cfg.StartIndex); err != nil {
		return &ErrSessionSetup{err}
	}

	return nil
}

// SetupWindow creates a new tmux window and executes the configured command
func (d *Terminal) SetupWindow(ctx context.Context, cfgIndex int) error {
	if cfgIndex != d.cfg.StartIndex {
		if err := d.tmux.NewWindow(ctx, cfgIndex); err != nil {
			return &ErrWindowSetup{err}
		}
	}

	d.cfg.Windows[cfgIndex].Cmd = append(d.cfg.Windows[cfgIndex].Cmd, TmuxEnterCmd)

	if err := d.tmux.SendKeys(ctx, cfgIndex); err != nil {
		return &ErrWindowSetup{err}
	}

	return nil
}
