package domain

import (
	"context"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
	"github.com/nxdir-s/gomux/internal/ports"
)

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
	exists, err := d.SessionExists(ctx)
	if err != nil {
		return err
	}

	if exists == tmux.SessionNotExists {
		if err := d.SetupSession(ctx); err != nil {
			return err
		}
	}

	if err := d.Attach(ctx); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) SessionExists(ctx context.Context) (int, error) {
	return d.service.SessionExists(ctx)
}

func (d *Tmux) Attach(ctx context.Context) error {
	return d.service.AttachSession(ctx)
}

func (d *Tmux) SetupSession(ctx context.Context) error {
	if err := d.service.NewSession(ctx, d.cfg.Session); err != nil {
		return err
	}

	for index := range d.cfg.Windows {
		if err := d.SetupWindow(ctx, index); err != nil {
			return err
		}
	}

	if err := d.service.SelectWindow(ctx, d.cfg.StartIndex); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) SetupWindow(ctx context.Context, cfgIndex int) error {
	if err := d.service.NewWindow(ctx, cfgIndex); err != nil {
		return err
	}

	if err := d.GoToProject(ctx, cfgIndex); err != nil {
		return err
	}

	if err := d.service.SendKeys(ctx, d.cfg.Windows[cfgIndex].Name, d.cfg.Windows[cfgIndex].Cmd); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) GoToProject(ctx context.Context, cfgIndex int) error {
	if err := d.service.SendKeys(ctx, d.cfg.Windows[cfgIndex].Name, "cd "+d.cfg.Project); err != nil {
		return err
	}

	return nil
}
