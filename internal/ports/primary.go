package ports

import "context"

type CLIPort interface {
	Start(ctx context.Context) error
}
