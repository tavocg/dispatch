//go:build !windows

package dispatch

import (
	"context"
	"os"
	"os/exec"
)

type UnixEscalator struct{}

func NewEscalator() Escalator { return UnixEscalator{} }

func (UnixEscalator) CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd {
	if os.Geteuid() == 0 {
		return exec.CommandContext(ctx, name, arg...)
	}

	if !isInteractive(ctx) {
		return nil
	}

	if p, err := exec.LookPath(SudoProg); err == nil {
		args := append([]string{"--", name}, arg...)
		cmd := exec.CommandContext(ctx, p, args...)
		cmd.Stdin = os.Stdin
		return cmd
	}

	if p, err := exec.LookPath(DoasProg); err == nil {
		args := append([]string{"--", name}, arg...)
		cmd := exec.CommandContext(ctx, p, args...)
		cmd.Stdin = os.Stdin
		return cmd
	}

	if p, err := exec.LookPath(PkexecProg); err == nil {
		args := append([]string{name}, arg...)
		cmd := exec.CommandContext(ctx, p, args...)
		cmd.Stdin = os.Stdin
		return cmd
	}

	return nil
}
