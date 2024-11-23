package primary

import (
	"context"

	"github.com/nxdir-s/gomux/internal/ports"
)

type CLIAdapter struct {
	tmux ports.Tmux
}

// NewCLIAdapter creates a CLIAdapter
func NewCLIAdapter(tmux ports.Tmux) (*CLIAdapter, error) {
	return &CLIAdapter{
		tmux: tmux,
	}, nil
}

// TmuxStart starts tmux
func (a *CLIAdapter) TmuxStart(ctx context.Context) error {
	if err := a.tmux.Start(ctx); err != nil {
		return err
	}

	return nil
}
