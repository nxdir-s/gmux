package domain

import (
	"context"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
	"github.com/nxdir-s/gomux/internal/ports"
)

type ErrStartTmux struct {
	err error
}

func (e *ErrStartTmux) Error() string {
	return "failed tmux setup: " + e.err.Error()
}

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

type ErrCDProject struct {
	err error
}

func (e *ErrCDProject) Error() string {
	return "failed to go to project directory: " + e.err.Error()
}

type Tmux struct {
	cfg     *entity.Config
	service ports.TmuxService
}

func NewTmux(config *entity.Config, service ports.TmuxService) (*Tmux, error) {
	return &Tmux{
		cfg:     config,
		service: service,
	}, nil
}

func (d *Tmux) Start(ctx context.Context) error {
	if exists := d.service.SessionExists(ctx); exists == tmux.SessionNotExists {
		if err := d.SetupSession(ctx); err != nil {
			return &ErrStartTmux{err}
		}
	}

	if err := d.service.AttachSession(ctx); err != nil {
		return &ErrStartTmux{err}
	}

	return nil
}

func (d *Tmux) SetupSession(ctx context.Context) error {
	if err := d.service.NewSession(ctx, d.cfg.Windows[d.cfg.StartIndex].Name); err != nil {
		return &ErrSessionSetup{err}
	}

	for index := range d.cfg.Windows {
		if err := d.SetupWindow(ctx, index); err != nil {
			return &ErrSessionSetup{err}
		}
	}

	if err := d.service.SelectWindow(ctx, d.cfg.StartIndex); err != nil {
		return &ErrSessionSetup{err}
	}

	return nil
}

func (d *Tmux) SetupWindow(ctx context.Context, cfgIndex int) error {
	if err := d.service.NewWindow(ctx, cfgIndex); err != nil {
		return &ErrWindowSetup{err}
	}

	if err := d.GoToProject(ctx, cfgIndex); err != nil {
		return &ErrWindowSetup{err}
	}

	if err := d.service.SendKeys(ctx, d.cfg.Windows[cfgIndex].Name, d.cfg.Windows[cfgIndex].Cmd...); err != nil {
		return &ErrWindowSetup{err}
	}

	return nil
}

func (d *Tmux) GoToProject(ctx context.Context, cfgIndex int) error {
	if err := d.service.SendKeys(ctx, d.cfg.Windows[cfgIndex].Name, "cd", d.cfg.Project); err != nil {
		return &ErrCDProject{err}
	}

	return nil
}
