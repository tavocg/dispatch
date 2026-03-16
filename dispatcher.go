// Package dispatch
package dispatch

import "os/exec"

type Dispatcher interface {
	WithStreamer(Streamer) Dispatcher
	WithEscalator(Escalator) Dispatcher
	WithPrivileged() Dispatcher
	WithInteractive() Dispatcher
	Command(name string, arg ...string) *exec.Cmd
}
