package secondary

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"testing"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/tmux"
	"github.com/nxdir-s/gomux/tests"
)

const (
	TmuxSessionName      string = "TmuxUnitTests"
	TmuxSessionExists    string = "HasSession should return 0 if a tmux session exists"
	TmuxSessionNotExists string = "HasSession should return 1 if a tmux session does not exists"
)

func TestHasSession(t *testing.T) {
	var cases = []struct {
		name      string
		cfg       *entity.Config
		expected  int
		shouldErr bool
	}{
		{
			name: TmuxSessionExists,
			cfg: &entity.Config{
				Session: TmuxSessionName,
			},
			expected:  tmux.SessionExists,
			shouldErr: false,
		},
		{
			name: TmuxSessionNotExists,
			cfg: &entity.Config{
				Session: TmuxSessionName,
			},
			expected:  tmux.SessionNotExists,
			shouldErr: true,
		},
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cmd := exec.CommandContext(ctx, tmux.Alias, string(tmux.HasSessionCmd), "-t", TmuxSessionName)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockCmd, err := tests.NewCommandMock(tc.cfg, cmd, tc.shouldErr)
			if err != nil {
				t.Errorf("failed to created CommandMock: %s", err.Error())
			}

			adapter, err := NewTmuxAdapter(tc.cfg, mockCmd)
			if err != nil {
				t.Errorf("failed to created TmuxAdapter: %s", err.Error())
			}

			if out := adapter.HasSession(ctx); out != tc.expected {
				t.Errorf("error got %d, want %d", out, tc.expected)
			}
		})
	}
}
