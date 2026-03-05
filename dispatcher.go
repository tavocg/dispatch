// Package dispatch
package dispatch

import (
	"context"
	"errors"
	"os"
	"os/exec"
)

type Dispatcher struct {
	p         *DispatcherParams
	escalator Escalator
}

type DispatcherParams struct {
	Ctx         context.Context
	Streamer    Streamer
	Privileged  bool
	Interactive bool
}

func NewDispatcher(params *DispatcherParams) *Dispatcher {
	return &Dispatcher{
		p:         params,
		escalator: NewEscalator(),
	}
}

func (d *Dispatcher) Run(name string, arg ...string) error {
	var cmd *exec.Cmd

	d.p.Ctx = withInteractive(d.p.Ctx, d.p.Interactive)

	if d.p.Privileged {
		if !d.p.Interactive {
			return errors.New("not interactive, cannot escalate privileges")
		}

		cmd = d.escalator.CommandContext(d.p.Ctx, name, arg...)

		if cmd == nil {
			return errors.New("could not escalate privileges")
		}
	} else {
		cmd = exec.CommandContext(d.p.Ctx, name, arg...)
	}

	if d.p.Streamer != nil {
		if w := d.p.Streamer.Stdout(); w != nil {
			cmd.Stdout = w
		}
		if w := d.p.Streamer.Stderr(); w != nil {
			cmd.Stderr = w
		}
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
