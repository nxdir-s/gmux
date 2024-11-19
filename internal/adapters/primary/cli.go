package primary

import "context"

type CLIAdapter struct{}

func NewCLIAdapter() (*CLIAdapter, error) {
	return &CLIAdapter{}, nil
}

func (a *CLIAdapter) Start(ctx context.Context) error {
	return nil
}
