//go:build windows

package dispatch

import (
	"context"
	"os/exec"
)

// This package is not supported on Windows.
type WindowsEscalator struct{}

func NewEscalator() Escalator { return WindowsEscalator{} }

func (WindowsEscalator) CommandContext(context.Context, string, ...string) *exec.Cmd { return nil }
