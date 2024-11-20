package secondary

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
)

const (
	EnterCmd string = "C-m"
)

type TmuxAdapter struct {
	cfg *entity.Config
}

func NewTmuxAdapter(ctx context.Context, config *entity.Config) (*TmuxAdapter, error) {
	return &TmuxAdapter{
		cfg: config,
	}, nil
}

func (a *TmuxAdapter) HasSession(ctx context.Context) int {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.HasSessionCmd), "-t", a.cfg.Session)

	fmt.Fprintf(os.Stdout, "checking for existing session '%s'\n", a.cfg.Session)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", string(output))

		return tmux.SessionNotExists
	}

	fmt.Fprintf(os.Stdout, "%s cmd output: %s\n", string(tmux.HasSessionCmd), string(output))

	return tmux.SessionExists
}

func (a *TmuxAdapter) NewSession(ctx context.Context, name string) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewSessionCmd), "-d", "-s", a.cfg.Session)

	fmt.Fprintf(os.Stdout, "creating new session named '%s'\n", a.cfg.Session)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed %s cmd, output: %s\n", string(tmux.NewSessionCmd), string(output))

		return err
	}

	fmt.Fprintf(os.Stdout, "%s cmd output: %s\n", string(tmux.NewSessionCmd), string(output))

	return nil
}

func (a *TmuxAdapter) AttachSession(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.AttachCmd), "-t", a.cfg.Session)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed %s cmd, output: %s\n", string(tmux.AttachCmd), string(output))

		return err
	}

	fmt.Fprintf(os.Stdout, "%s cmd output: %s\n", string(tmux.AttachCmd), string(output))

	return nil
}

func (a *TmuxAdapter) SendKeys(ctx context.Context, name string, args ...string) error {
	cmdArgs := []string{string(tmux.SendKeysCmd), "-t", a.cfg.Session + ":" + name}
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.CommandContext(ctx, tmux.Alias, cmdArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed %s cmd, output: %s\n", string(tmux.SendKeysCmd), string(output))

		return err
	}

	fmt.Fprintf(os.Stdout, "%s cmd output: %s\n", string(tmux.SendKeysCmd), string(output))

	return nil
}

func (a *TmuxAdapter) NewWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.NewWindowCmd), "-t", a.cfg.Session, "-n", a.cfg.Windows[cfgIndex].Name)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed %s cmd, output: %s\n", string(tmux.NewWindowCmd), string(output))

		return err
	}

	fmt.Fprintf(os.Stdout, "%s cmd output: %s\n", string(tmux.NewWindowCmd), string(output))

	return nil
}

func (a *TmuxAdapter) SelectWindow(ctx context.Context, cfgIndex int) error {
	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.SelectWindowCmd), "-t", a.cfg.Session+":"+a.cfg.Windows[cfgIndex].Name)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed %s cmd, output: %s\n", string(tmux.SelectWindowCmd), string(output))

		return err
	}

	fmt.Fprintf(os.Stdout, "%s cmd output: %s\n", string(tmux.SelectWindowCmd), string(output))

	return nil
}
