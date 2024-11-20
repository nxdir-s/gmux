package ports

import (
	"context"
)

type Tmux interface {
	Start(ctx context.Context) error
}

type TmuxService interface {
	SessionExists(ctx context.Context) (int, error)
	NewSession(ctx context.Context, name string) error
	AttachSession(ctx context.Context) error
	SendKeys(ctx context.Context, name string, keyCmd string) error
	NewWindow(ctx context.Context, cfgIndex int) error
	SelectWindow(ctx context.Context, cfgIndex int) error
}
