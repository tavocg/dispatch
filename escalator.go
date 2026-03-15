package dispatch

import (
	"context"
	"os/exec"
)

type Escalator interface {
	IsPrivilegedUser() bool
	CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd
}
