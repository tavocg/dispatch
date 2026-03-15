//go:build !windows

package dispatch

import (
	"context"
	"errors"
	"os/exec"
)

type UnixDispatcher struct {
	p *UnixDispatcherParams
}

type UnixDispatcherParams struct {
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

	return &UnixDispatcher{
		p: &UnixDispatcherParams{
			Ctx:       ctx,
			Streamer:  NewDefaultStreamer(),
			Escalator: NewDefaultEscalator(),
		},
	}, nil
}

func (d *UnixDispatcher) WithStreamer(s Streamer) Dispatcher {
	if s != nil {
		d.p.Streamer = s
	}
	return d
}

func (d *UnixDispatcher) WithEscalator(e Escalator) Dispatcher {
	if e != nil {
		d.p.Escalator = e
	}
	return d
}

func (d *UnixDispatcher) WithPrivileged() Dispatcher {
	d.p.Interactive = true
	d.p.Privileged = true
	return d
}

func (d *UnixDispatcher) WithInteractive() Dispatcher {
	d.p.Interactive = true
	return d
}

func (d *UnixDispatcher) Run(name string, arg ...string) error {
	var cmd *exec.Cmd
	ctx := d.p.Ctx

	cmd = exec.CommandContext(ctx, name, arg...)

	if d.p.Privileged && !d.p.Escalator.IsPrivilegedUser() {
		cmd = d.p.Escalator.CommandContext(ctx, name, arg...)
	}

	if cmd == nil {
		return errors.New("nil cmd")
	}

	cmd.Stdout = d.p.Streamer.Stdout()
	cmd.Stderr = d.p.Streamer.Stderr()

	return cmd.Run()
}
