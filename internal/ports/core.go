package ports

import (
	"context"
)

type Tmux interface {
	Start(ctx context.Context) error
	SetupSession(ctx context.Context) error
	SetupWindow(ctx context.Context, cfgIndex int) error
}

type TmuxService interface {
	SessionExists(ctx context.Context) int
	NewSession(ctx context.Context, name string) error
	NewWindow(ctx context.Context, cfgIndex int) error
	SelectWindow(ctx context.Context, cfgIndex int) error
	SendKeys(ctx context.Context, name string, args ...string) error
	AttachSession(ctx context.Context) error
}
