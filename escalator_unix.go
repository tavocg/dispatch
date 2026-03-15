//go:build !windows

package dispatch

import (
	"context"
	"os"
	"os/exec"
)

const (
	sudoProg   = "sudo"
	doasProg   = "doas"
	pkexecProg = "pkexec"
)

type UnixEscalator struct{}

type DefaultEscalator = UnixEscalator

func NewDefaultEscalator() Escalator { return &DefaultEscalator{} }

func NewEscalator() Escalator { return NewDefaultEscalator() }

func NewUnixEscalator() Escalator { return &UnixEscalator{} }

func (UnixEscalator) IsPrivilegedUser() bool {
	return os.Geteuid() == 0
}

func (u *UnixEscalator) CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd {
	if u.IsPrivilegedUser() {
		return exec.CommandContext(ctx, name, arg...)
	}

	if p, err := exec.LookPath(sudoProg); err == nil {
		args := append([]string{"--", name}, arg...)
		cmd := exec.CommandContext(ctx, p, args...)
		cmd.Stdin = os.Stdin
		return cmd
	}

	if p, err := exec.LookPath(doasProg); err == nil {
		args := append([]string{"--", name}, arg...)
		cmd := exec.CommandContext(ctx, p, args...)
		cmd.Stdin = os.Stdin
		return cmd
	}

	if p, err := exec.LookPath(pkexecProg); err == nil {
		args := append([]string{name}, arg...)
		cmd := exec.CommandContext(ctx, p, args...)
		cmd.Stdin = os.Stdin
		return cmd
	}

	return nil
}
