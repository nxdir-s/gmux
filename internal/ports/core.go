package ports

import (
	"context"
)

type Terminal interface {
	StartTmux(ctx context.Context) error
	SetupSession(ctx context.Context) error
	SetupWindow(ctx context.Context, cfgIndex int) error
}

type TerminalService interface {
	TmuxSessionExists(ctx context.Context, session string) int
	NewTmuxSession(ctx context.Context, name string) error
	NewTmuxWindow(ctx context.Context, session string, name string) error
	TmuxSelectWindow(ctx context.Context, session string, window string) error
	TmuxSendKeys(ctx context.Context, cmd []string, session string, window string) error
	TmuxAttachSession(ctx context.Context, session string) error
}
