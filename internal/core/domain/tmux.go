package domain

import "context"

type Tmux struct{}

func NewTmux() (*Tmux, error) {
	return &Tmux{}, nil
}

func (d *Tmux) Start(ctx context.Context) error {
	return nil
}
