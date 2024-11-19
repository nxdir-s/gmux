package ports

import "context"

type TmuxPort interface {
	HasSession(ctx context.Context) (int, error)
	NewSession(ctx context.Context) error
	AttachSession(ctx context.Context) error
	SendKeys(ctx context.Context) error
	NewWindow(ctx context.Context) error
	SelectWindow(ctx context.Context) error
	SetOption(ctx context.Context) error
	SetWindowOpt(ctx context.Context) error
}
