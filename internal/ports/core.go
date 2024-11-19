package ports

import "context"

type Tmux interface {
	Start(ctx context.Context) error
}
