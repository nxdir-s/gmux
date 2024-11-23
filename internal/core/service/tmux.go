package service

import (
	"context"

	"github.com/nxdir-s/gomux/internal/ports"
)

type TmuxService struct {
	tmux ports.TmuxPort
}

func NewTmuxService(adapter ports.TmuxPort) (*TmuxService, error) {
	return &TmuxService{
		tmux: adapter,
	}, nil
}

// SessionExists checks if a tmux session exists
func (s *TmuxService) SessionExists(ctx context.Context) int {
	return s.tmux.HasSession(ctx)
}

// NewSession creates a new session
func (s *TmuxService) NewSession(ctx context.Context, name string) error {
	return s.tmux.NewSession(ctx, name)
}

// NewWindow creates a new window
func (s *TmuxService) NewWindow(ctx context.Context, cfgIndex int) error {
	return s.tmux.NewWindow(ctx, cfgIndex)
}

// SelectWindow selects a tmux window
func (s *TmuxService) SelectWindow(ctx context.Context, cfgIndex int) error {
	return s.tmux.SelectWindow(ctx, cfgIndex)
}

// SendKeys executes the send-keys tmux command
func (s *TmuxService) SendKeys(ctx context.Context, cfgIndex int) error {
	return s.tmux.SendKeys(ctx, cfgIndex)
}

// AttachSession attempts to attach to a tmux session
func (s *TmuxService) AttachSession(ctx context.Context) error {
	return s.tmux.AttachSession(ctx)
}
