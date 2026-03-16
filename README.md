# dispatch

Small Go package to run system commands with optional privilege escalation.

## Install

```bash
go get github.com/tavocg/dispatch
```

## Example usage

```go
package main

import (
	"context"
	"log"

	"github.com/tavocg/dispatch"
)

func main() {
	d, err := dispatch.NewDispatcher(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Runs via sudo/doas/pkexec on Unix when needed.
	cmd := d.WithPrivileged().WithInteractive().Command("whoami")
	if cmd == nil {
		log.Fatal("failed to create command")
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
```
