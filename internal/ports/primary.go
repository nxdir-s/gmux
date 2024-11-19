package ports

import "context"

type CLIPort interface {
	TmuxStart(ctx context.Context) error
}
