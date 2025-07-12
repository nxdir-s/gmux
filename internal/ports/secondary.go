package ports

import (
	"context"
)

type Tmux interface {
	SessionExists(ctx context.Context, session string) int
	NewSession(ctx context.Context, name string) error
	AttachSession(ctx context.Context, session string) error
	SendKeys(ctx context.Context, cmd []string, session string, window string) error
	NewWindow(ctx context.Context, session string, name string) error
	SelectWindow(ctx context.Context, session string, window string) error
}
