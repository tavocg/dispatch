//go:build !windows

package dispatch

import (
	"context"
	"os"
	"os/exec"
	"testing"
)

func TestIntegrationRunPrivilegedWhoamiUnix(t *testing.T) {
	if os.Getenv("INTEGRATION") != "1" {
		t.Skip("set INTEGRATION=1 to run integration tests")
	}

	if _, err := exec.LookPath("whoami"); err != nil {
		t.Skip("whoami command is not available")
	}

	if os.Geteuid() == 0 {
		t.Skip("running as root does not require privilege escalation prompt")
	}

	_, hasSudo := exec.LookPath(SudoProg)
	_, hasDoas := exec.LookPath(DoasProg)
	_, hasPkexec := exec.LookPath(PkexecProg)

	if hasSudo != nil && hasDoas != nil && hasPkexec != nil {
		t.Skip("one of sudo, doas, or pkexec is required for this integration test")
	}

	if hasSudo == nil {
		// Invalidate cached credentials so sudo can prompt again when needed.
		_ = exec.Command(SudoProg, "-k").Run()
	}

	if hasDoas == nil {
		// Best effort: clear persisted doas auth on implementations that support -L.
		_ = exec.Command(DoasProg, "-L").Run()
	}

	d := NewDispatcher(&DispatcherParams{
		Ctx: context.Background(),
	})
	d.WithPrivileged().WithInteractive()

	if err := d.Run("whoami"); err != nil {
		t.Fatalf("privileged whoami failed: %v", err)
	}
}
