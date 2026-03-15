// Package dispatch
package dispatch

type Dispatcher interface {
	WithStreamer(Streamer) Dispatcher
	WithEscalator(Escalator) Dispatcher
	WithPrivileged() Dispatcher
	WithInteractive() Dispatcher
	Run(name string, arg ...string) error
}
