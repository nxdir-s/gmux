package ports

import (
	"context"

	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
)

type Tmux interface {
	Start(ctx context.Context) error
}

type TmuxService interface {
	SessionExists(ctx context.Context) (int, error)
	NewSession(ctx context.Context) error
	AttachSession(ctx context.Context) error
	SendKeys(ctx context.Context, window tmux.Window, keyCmd string) error
	NewWindow(ctx context.Context, window tmux.Window) error
	SelectWindow(ctx context.Context, window tmux.Window) error
}
