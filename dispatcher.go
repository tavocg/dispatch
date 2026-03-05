// Package dispatch
package dispatch

import (
	"context"
	"errors"
	"os/exec"
)

type Dispatcher struct {
	p         *DispatcherParams
	escalator Escalator
}

type DispatcherParams struct {
	Ctx      context.Context
	Streamer Streamer
}

func NewDispatcher(params *DispatcherParams) *Dispatcher {
	if params == nil {
		params = &DispatcherParams{}
	}

	if params.Ctx == nil {
		params.Ctx = context.Background()
	}

	if params.Streamer == nil {
		params.Streamer = NewDefaultStreamer()
	}

	return &Dispatcher{
		p:         params,
		escalator: NewEscalator(),
	}
}

func (d *Dispatcher) WithPrivileged() *Dispatcher {
	d.p.Ctx = withPrivileged(d.p.Ctx)
	return d
}

func (d *Dispatcher) WithInteractive() *Dispatcher {
	d.p.Ctx = withInteractive(d.p.Ctx)
	return d
}

func (d *Dispatcher) Run(name string, arg ...string) error {
	var cmd *exec.Cmd
	ctx := d.p.Ctx

	if isPrivileged(ctx) {
		cmd = d.escalator.CommandContext(ctx, name, arg...)
		if cmd == nil {
			return errors.New("could not escalate privileges")
		}
	} else {
		cmd = exec.CommandContext(ctx, name, arg...)
	}

	cmd.Stdout = d.p.Streamer.Stdout()
	cmd.Stderr = d.p.Streamer.Stderr()

	return cmd.Run()
}
