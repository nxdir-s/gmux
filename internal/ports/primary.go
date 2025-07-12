package ports

import "context"

type CLI interface {
	StartTmux(ctx context.Context) error
}
