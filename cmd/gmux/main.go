package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/nxdir-s/adapters"
	"github.com/nxdir-s/gmux/internal/adapters/primary"
	"github.com/nxdir-s/gmux/internal/core/domain"
	"github.com/nxdir-s/gmux/internal/core/valobj"
	"github.com/nxdir-s/gmux/internal/ports"
)

const CfgFileName string = ".gmux.toml"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// domain orchestrators
	var terminal ports.Terminal

	// secondary adapters
	var tmux ports.Tmux

	// primary adapter
	var cli ports.CLI

	toml := adapters.NewTomlAdapter[valobj.Config]()

	if err := toml.LoadConfig(CfgFileName); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		os.Exit(1)
	}

	tmux = adapters.NewTmuxAdapter(adapters.NewCmdAdapter())
	terminal = domain.NewTerminal(&toml.Cfg, tmux)
	cli = primary.NewCLIAdapter(terminal)

	if err := cli.StartTmux(ctx); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		os.Exit(1)
	}
}
