package domain

import (
	"context"

	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
	"github.com/nxdir-s/gomux/internal/core/valobj"
	"github.com/nxdir-s/gomux/internal/ports"
)

type Tmux struct {
	cfg     *valobj.Config
	service ports.TmuxService
}

func NewTmux(config *valobj.Config, service ports.TmuxService) (*Tmux, error) {
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

func (d *Tmux) Attach(ctx context.Context) error {
	return d.service.AttachSession(ctx)
}

func (d *Tmux) SessionExists(ctx context.Context) (int, error) {
	return d.service.SessionExists(ctx)
}

func (d *Tmux) GoToProject(ctx context.Context, window tmux.Window) error {
	if err := d.service.SendKeys(ctx, window, "cd "+d.cfg.ProjectDir); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) SetupSession(ctx context.Context) error {
	if err := d.service.NewSession(ctx); err != nil {
		return err
	}

	if err := d.SetupEditor(ctx, tmux.Editor); err != nil {
		return err
	}

	if err := d.SetupDocker(ctx, tmux.Docker); err != nil {
		return err
	}

	if err := d.SetupDatabase(ctx, tmux.Database); err != nil {
		return err
	}

	if err := d.service.SelectWindow(ctx, tmux.Editor); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) SetupWindow(ctx context.Context, window tmux.Window, cmd string) error {
	if err := d.service.NewWindow(ctx, window); err != nil {
		return err
	}

	if err := d.GoToProject(ctx, window); err != nil {
		return err
	}

	if err := d.service.SendKeys(ctx, window, cmd); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) SetupDocker(ctx context.Context, window tmux.Window) error {
	if err := d.service.NewWindow(ctx, window); err != nil {
		return err
	}

	if err := d.GoToProject(ctx, window); err != nil {
		return err
	}

	if err := d.service.SendKeys(ctx, window, d.cfg.DockerCmd); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) SetupEditor(ctx context.Context, window tmux.Window) error {
	if err := d.GoToProject(ctx, window); err != nil {
		return err
	}

	if err := d.service.SendKeys(ctx, window, d.cfg.EditorCmd); err != nil {
		return err
	}

	return nil
}

func (d *Tmux) SetupDatabase(ctx context.Context, window tmux.Window) error {
	if err := d.service.NewWindow(ctx, window); err != nil {
		return err
	}

	if err := d.service.SendKeys(ctx, window, d.cfg.DatabaseCmd); err != nil {
		return err
	}

	return nil
}
