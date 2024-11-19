package service

import (
	"context"

	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
	"github.com/nxdir-s/gomux/internal/ports"
)

type TmuxService struct {
	adapter ports.TmuxPort
}

func NewTmuxService() (*TmuxService, error) {
	return &TmuxService{}, nil
}

func (s *TmuxService) SessionExists(ctx context.Context) (int, error) {
	return s.adapter.HasSession(ctx)
}

func (s *TmuxService) NewSession(ctx context.Context) error {
	return s.adapter.NewSession(ctx)
}

func (s *TmuxService) AttachSession(ctx context.Context) error {
	return s.adapter.AttachSession(ctx)
}

func (s *TmuxService) SendKeys(ctx context.Context, window tmux.Window, keyCmd string) error {
	return s.adapter.SendKeys(ctx, window, keyCmd)
}

func (s *TmuxService) NewWindow(ctx context.Context, window tmux.Window) error {
	return s.adapter.NewWindow(ctx, window)
}

func (s *TmuxService) SelectWindow(ctx context.Context, window tmux.Window) error {
	return s.adapter.SelectWindow(ctx, window)
}
