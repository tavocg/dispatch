//go:build !windows

package dispatch

import (
	"context"
	"os"
	"os/exec"
	"testing"
)

func TestIntegrationRunPrivilegedWhoami(t *testing.T) {
	if os.Getenv("INTEGRATION") != "1" {
		t.Skip("set INTEGRATION=1 to run integration tests")
	}

	if _, err := exec.LookPath("whoami"); err != nil {
		t.Skip("whoami command is not available")
	}

	if os.Geteuid() == 0 {
		t.Skip("running as root does not require sudo password prompt")
	}

	if _, err := exec.LookPath(SudoProg); err != nil {
		t.Skip("sudo is required for this integration test")
	}

	// Invalidate cached credentials so sudo can prompt again when needed.
	_ = exec.Command(SudoProg, "-k").Run()

	d := NewDispatcher(&DispatcherParams{
		Ctx: context.Background(),
	})
	d.WithPrivileged().WithInteractive()

	if err := d.Run("whoami"); err != nil {
		t.Fatalf("privileged whoami failed: %v", err)
	}
}
