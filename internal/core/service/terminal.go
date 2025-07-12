package service

import (
	"context"

	"github.com/nxdir-s/gmux/internal/ports"
)

type Terminal struct {
	tmux ports.Tmux
}

func NewTerminal(adapter ports.Tmux) *Terminal {
	return &Terminal{
		tmux: adapter,
	}
}

// TmuxSessionExists checks if a tmux session exists
func (s *Terminal) TmuxSessionExists(ctx context.Context, session string) int {
	return s.tmux.HasSession(ctx, session)
}

// NewTmuxSession creates a new session
func (s *Terminal) NewTmuxSession(ctx context.Context, name string) error {
	return s.tmux.NewSession(ctx, name)
}

// NewTmuxWindow creates a new window
func (s *Terminal) NewTmuxWindow(ctx context.Context, session string, name string) error {
	return s.tmux.NewWindow(ctx, session, name)
}

// SelectTmuxWindow selects a tmux window
func (s *Terminal) TmuxSelectWindow(ctx context.Context, session string, window string) error {
	return s.tmux.SelectWindow(ctx, session, window)
}

// TmuxSendKeys executes the send-keys tmux command
func (s *Terminal) TmuxSendKeys(ctx context.Context, cmd []string, session string, window string) error {
	return s.tmux.SendKeys(ctx, cmd, session, window)
}

// AttachSession attempts to attach to a tmux session
func (s *Terminal) TmuxAttachSession(ctx context.Context, session string) error {
	return s.tmux.AttachSession(ctx, session)
}
