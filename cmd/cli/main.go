package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/nxdir-s/gomux/internal/adapters/primary"
	"github.com/nxdir-s/gomux/internal/adapters/secondary"
	"github.com/nxdir-s/gomux/internal/core/domain"
	"github.com/nxdir-s/gomux/internal/ports"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// domains
	var tmux ports.Tmux

	// secondary adapters
	var tmuxAdapter ports.TmuxPort
	var config ports.ConfigPort

	// primary adapters
	var cli ports.CLIPort

	config, err := secondary.NewTomlAdapter()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating config adapter: %s\n", err.Error())
		os.Exit(1)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error loading config: %s\n", err.Error())
		os.Exit(1)
	}

	tmuxAdapter, err = secondary.NewTmuxAdapter(cfg)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating tmux adapter: %s\n", err.Error())
		os.Exit(1)
	}

	tmux, err = domain.NewTmux(tmuxAdapter)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating tmux orchestrator: %s\n", err.Error())
		os.Exit(1)
	}

	cli, err = primary.NewCLIAdapter(tmux)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating cli adapter: %s\n", err.Error())
		os.Exit(1)
	}

	if err := cli.TmuxStart(ctx); err != nil {
		fmt.Fprintf(os.Stdout, "error starting tmux: %s\n", err.Error())
		os.Exit(1)
	}
}
