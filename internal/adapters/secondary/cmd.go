package secondary

import (
	"context"
	"os/exec"
)

type CommandAdapter struct{}

func NewCommandAdapter(ctx context.Context) (*CommandAdapter, error) {
	return &CommandAdapter{}, nil
}

func (a *CommandAdapter) Exec(ctx context.Context, cmd *exec.Cmd) ([]byte, error) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
