package dispatch

import (
	"io"
	"os"
)

type Streamer interface {
	Stdout() io.Writer
	Stderr() io.Writer
}

type DefaultStreamer struct{}

func NewDefaultStreamer() DefaultStreamer {
	return DefaultStreamer{}
}

func (DefaultStreamer) Stdout() io.Writer {
	return os.Stdout
}

func (DefaultStreamer) Stderr() io.Writer {
	return os.Stderr
}
