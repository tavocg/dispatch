// Package dispatch is not supported on Windows.
//go:build windows

package dispatch

import (
	"context"
	"os/exec"
)

type WindowsEscalator struct{}

type DefaultEscalator = WindowsEscalator

func NewDefaultEscalator() Escalator { return DefaultEscalator{} }

func NewEscalator() Escalator { return NewDefaultEscalator() }

func (WindowsEscalator) IsPrivilegedUser() bool { return false }

func (WindowsEscalator) CommandContext(context.Context, string, ...string) *exec.Cmd { return nil }
