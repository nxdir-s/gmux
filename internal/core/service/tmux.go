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

func (s *TmuxService) SessionExists(ctx context.Context) int {
	return s.tmux.HasSession(ctx)
}

func (s *TmuxService) NewSession(ctx context.Context, name string) error {
	return s.tmux.NewSession(ctx, name)
}

func (s *TmuxService) NewWindow(ctx context.Context, cfgIndex int) error {
	return s.tmux.NewWindow(ctx, cfgIndex)
}

func (s *TmuxService) SelectWindow(ctx context.Context, cfgIndex int) error {
	return s.tmux.SelectWindow(ctx, cfgIndex)
}

func (s *TmuxService) SendKeys(ctx context.Context, name string, args ...string) error {
	return s.tmux.SendKeys(ctx, name, args...)
}

func (s *TmuxService) AttachSession(ctx context.Context) error {
	return s.tmux.AttachSession(ctx)
}
