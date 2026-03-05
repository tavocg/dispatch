package dispatch

import (
	"io"
)

type Streamer interface {
	Stdout() io.Writer
	Stderr() io.Writer
}
