package dispatch

import (
	"bytes"
	"context"
	"io"
	"os/exec"
	"strings"
	"testing"
)

type passthroughEscalator struct {
	called bool
	name   string
	args   []string
}

func (e *passthroughEscalator) CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd {
	e.called = true
	e.name = name
	e.args = append([]string(nil), arg...)
	return exec.CommandContext(ctx, name, arg...)
}

type bufferStreamer struct {
	stdout io.Writer
	stderr io.Writer
}

func (s *bufferStreamer) Stdout() io.Writer { return s.stdout }
func (s *bufferStreamer) Stderr() io.Writer { return s.stderr }

func TestDispatcherRunPrivilegedEscalatesWhoami(t *testing.T) {
	if _, err := exec.LookPath("whoami"); err != nil {
		t.Skip("whoami command is not available")
	}

	var out bytes.Buffer
	var errOut bytes.Buffer

	d := NewDispatcher(&DispatcherParams{
		Ctx:      context.Background(),
		Streamer: &bufferStreamer{stdout: &out, stderr: &errOut},
	})
	d.WithPrivileged().WithInteractive()

	escalator := &passthroughEscalator{}
	d.escalator = escalator

	if err := d.Run("whoami"); err != nil {
		t.Fatalf("run failed: %v (stderr=%q)", err, errOut.String())
	}

	if !escalator.called {
		t.Fatal("expected escalator to be called")
	}

	if escalator.name != "whoami" {
		t.Fatalf("expected command %q, got %q", "whoami", escalator.name)
	}

	if strings.TrimSpace(out.String()) == "" {
		t.Fatal("expected non-empty whoami output")
	}
}
