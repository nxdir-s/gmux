package ports

import (
	"context"

	"github.com/nxdir-s/gomux/internal/core/valobj"
)

type TmuxPort interface {
	HasSession(ctx context.Context) (int, error)
	NewSession(ctx context.Context) error
	AttachSession(ctx context.Context) error
	SendKeys(ctx context.Context, windowName string, keyCmd string) error
	NewWindow(ctx context.Context, name string) error
	SelectWindow(ctx context.Context) error
	SetOption(ctx context.Context) error
	SetWindowOpt(ctx context.Context) error
}

type ConfigPort interface {
	LoadConfig() (*valobj.Config, error)
}
