package tests

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"github.com/nxdir-s/gomux/internal/core/entity"
)

type ErrMockExec struct{}

func (e *ErrMockExec) Error() string {
	return "error executing command"
}

type ErrCmdArgs struct{}

func (e *ErrCmdArgs) Error() string {
	return "command arguments dont match"
}

type CommandMock struct {
	cfg       *entity.Config
	cmd       *exec.Cmd
	shouldErr bool
}

func NewCommandMock(config *entity.Config, cmd *exec.Cmd, shouldErr bool) (*CommandMock, error) {
	return &CommandMock{
		cfg:       config,
		cmd:       cmd,
		shouldErr: shouldErr,
	}, nil
}

func (a *CommandMock) Exec(ctx context.Context, cmd *exec.Cmd) ([]byte, error) {
	if !reflect.DeepEqual(a.cmd.Args, cmd.Args) {
		fmt.Fprintf(os.Stdout, "CommandMock: arguments are different: %+v %+v", a.cmd.Args, cmd.Args)
		return nil, &ErrCmdArgs{}
	}

	switch a.shouldErr {
	case true:
		return nil, &ErrMockExec{}
	case false:
		return make([]byte, 0), nil
	default:
		return make([]byte, 0), nil
	}
}
