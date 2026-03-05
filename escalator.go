package dispatch

import (
	"context"
	"os/exec"
)

const (
	SudoProg   = "sudo"
	DoasProg   = "doas"
	PkexecProg = "pkexec"
)

type Escalator interface {
	CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd
}
