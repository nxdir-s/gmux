package secondary

import "context"

type TmuxAdapter struct{}

func NewTmuxAdapter() (*TmuxAdapter, error) {
	return &TmuxAdapter{}, nil
}

func (a *TmuxAdapter) HasSession(ctx context.Context) (int, error) {
	return 1, nil
}

func (a *TmuxAdapter) NewSession(ctx context.Context) error {
	return nil
}

func (a *TmuxAdapter) AttachSession(ctx context.Context) error {
	return nil
}

func (a *TmuxAdapter) SendKeys(ctx context.Context) error {
	return nil
}

func (a *TmuxAdapter) NewWindow(ctx context.Context) error {
	return nil
}

func (a *TmuxAdapter) SelectWindow(ctx context.Context) error {
	return nil
}

func (a *TmuxAdapter) SetOption(ctx context.Context) error {
	return nil
}

func (a *TmuxAdapter) SetWindowOpt(ctx context.Context) error {
	return nil
}
