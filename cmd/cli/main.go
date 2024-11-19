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

	var tmux ports.Tmux
	var adapter ports.TmuxPort
	var cli ports.CLIPort

	adapter, err := secondary.NewTmuxAdapter()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating tmux adapter: %s\n", err.Error())
		os.Exit(1)
	}

	tmux, err = domain.NewTmux(adapter)
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
