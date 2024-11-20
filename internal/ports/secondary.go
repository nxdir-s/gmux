package ports

import (
	"context"

	"github.com/nxdir-s/gomux/internal/core/entity"
)

type TmuxPort interface {
	HasSession(ctx context.Context) int
	NewSession(ctx context.Context, name string) error
	AttachSession(ctx context.Context) error
	SendKeys(ctx context.Context, name string, args ...string) error
	NewWindow(ctx context.Context, cfgIndex int) error
	SelectWindow(ctx context.Context, cfgIndex int) error

	// SetOption(ctx context.Context) error
	// SetWindowOpt(ctx context.Context) error
}

type ConfigPort interface {
	LoadConfig() (*entity.Config, error)
}
