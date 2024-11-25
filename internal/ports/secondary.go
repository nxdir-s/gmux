package ports

import (
	"context"
	"os/exec"

	"github.com/nxdir-s/gomux/internal/core/entity"
)

type ConfigPort interface {
	LoadConfig() (*entity.Config, error)
}

type TmuxPort interface {
	HasSession(ctx context.Context) int
	NewSession(ctx context.Context, name string) error
	AttachSession(ctx context.Context) error
	SendKeys(ctx context.Context, cfgIndex int) error
	NewWindow(ctx context.Context, cfgIndex int) error
	SelectWindow(ctx context.Context, cfgIndex int) error
}

type CommandPort interface {
	Exec(ctx context.Context, cmd *exec.Cmd) ([]byte, error)
}
