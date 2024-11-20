package primary

import (
	"context"

	"github.com/nxdir-s/gomux/internal/ports"
)

type CLIAdapter struct {
	tmux ports.Tmux
}

func NewCLIAdapter(tmux ports.Tmux) (*CLIAdapter, error) {
	return &CLIAdapter{
		tmux: tmux,
	}, nil
}

func (a *CLIAdapter) TmuxStart(ctx context.Context) error {
	if err := a.tmux.Start(ctx); err != nil {
		return err
	}

	return nil
}
