// Package dispatch is not supported on Windows.
//go:build windows

package dispatch

import (
	"context"
	"errors"
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
			Escalator: NewEscalator(),
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
	d.p.Privileged = true
	return d
}

func (d *WindowsDispatcher) WithInteractive() Dispatcher {
	d.p.Interactive = true
	return d
}

func (d *WindowsDispatcher) Run(string, ...string) error {
	return errors.New("TODO: Windows dispatcher is not implemented")
}
