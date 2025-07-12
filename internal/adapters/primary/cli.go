package primary

import (
	"context"

	"github.com/nxdir-s/gmux/internal/ports"
)

type ErrStartTmux struct {
	err error
}

func (e *ErrStartTmux) Error() string {
	return "failed to start tmux: " + e.err.Error()
}

type CLIAdapter struct {
	terminal ports.Terminal
}

// NewCLIAdapter creates a CLIAdapter
func NewCLIAdapter(domain ports.Terminal) *CLIAdapter {
	return &CLIAdapter{
		terminal: domain,
	}
}

// StartTmux starts tmux in the terminal
func (a *CLIAdapter) StartTmux(ctx context.Context) error {
	if err := a.terminal.StartTmux(ctx); err != nil {
		return &ErrStartTmux{err}
	}

	return nil
}
