package tests

import (
	"bytes"
	"context"
	"fmt"
	"io"
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

func (a *CommandMock) Exec(ctx context.Context, cmd *exec.Cmd) (io.Reader, error) {
	if !reflect.DeepEqual(a.cmd.Args, cmd.Args) {
		fmt.Fprintf(os.Stdout, "CommandMock: arguments are different: %+v %+v", a.cmd.Args, cmd.Args)
		return nil, &ErrCmdArgs{}
	}

	switch a.shouldErr {
	case true:
		return bytes.NewReader([]byte("")), &ErrMockExec{}
	case false:
		return bytes.NewReader([]byte("")), nil
	default:
		return bytes.NewReader([]byte("")), nil
	}
}
