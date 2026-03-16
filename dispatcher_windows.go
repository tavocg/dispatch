// Package dispatch is not supported on Windows.
//go:build windows

package dispatch

import (
	"context"
	"errors"
	"os/exec"
)

type WindowsDispatcher struct {
	p *WindowsDispatcherParams
}

type WindowsDispatcherParams struct {
	Ctx         context.Context
	Streamer    Streamer
	Escalator   Escalator
	Privileged  bool
	Interactive bool
}

func NewDispatcher(ctx context.Context) (Dispatcher, error) {
	if ctx == nil {
		return nil, errors.New("nil ctx")
	}

	return &WindowsDispatcher{
		p: &WindowsDispatcherParams{
			Ctx:       ctx,
			Streamer:  NewDefaultStreamer(),
			Escalator: NewDefaultEscalator(),
		},
	}, nil
}

func (d *WindowsDispatcher) WithStreamer(s Streamer) Dispatcher {
	if s != nil {
		d.p.Streamer = s
	}
	return d
}

func (d *WindowsDispatcher) WithEscalator(e Escalator) Dispatcher {
	if e != nil {
		d.p.Escalator = e
	}
	return d
}

func (d *WindowsDispatcher) WithPrivileged() Dispatcher {
	d.p.Interactive = true
	d.p.Privileged = true
	return d
}

func (d *WindowsDispatcher) WithInteractive() Dispatcher {
	d.p.Interactive = true
	return d
}

func (d *WindowsDispatcher) Command(name string, arg ...string) *exec.Cmd {
	ctx := d.p.Ctx

	cmd := exec.CommandContext(ctx, name, arg...)
	if d.p.Privileged && !d.p.Escalator.IsPrivilegedUser() {
		cmd = d.p.Escalator.CommandContext(ctx, name, arg...)
	}

	if cmd == nil {
		return nil
	}

	cmd.Stdout = d.p.Streamer.Stdout()
	cmd.Stderr = d.p.Streamer.Stderr()

	return cmd
}
