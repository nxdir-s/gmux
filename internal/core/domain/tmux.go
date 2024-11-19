package domain

import (
	"context"

	"github.com/nxdir-s/gomux/internal/ports"
)

type Tmux struct {
	adapter ports.TmuxPort
}

func NewTmux(adapter ports.TmuxPort) (*Tmux, error) {
	return &Tmux{
		adapter: adapter,
	}, nil
}

func (d *Tmux) Start(ctx context.Context) error {
	return nil
}
