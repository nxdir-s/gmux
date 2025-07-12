package ports

import (
	"context"
)

type Terminal interface {
	StartTmux(ctx context.Context) error
	SetupSession(ctx context.Context) error
	SetupWindow(ctx context.Context, cfgIndex int) error
}
