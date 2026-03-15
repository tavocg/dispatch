// Package dispatch is not supported on Windows.
//go:build windows

package dispatch

import (
	"context"
	"os/exec"
)

type WindowsEscalator struct{}

func NewEscalator() Escalator { return WindowsEscalator{} }

func (WindowsEscalator) IsPrivilegedUser() bool { return false }

func (WindowsEscalator) CommandContext(context.Context, string, ...string) *exec.Cmd { return nil }
