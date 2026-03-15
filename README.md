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
	if err := d.WithPrivileged().WithInteractive().Run("whoami"); err != nil {
		log.Fatal(err)
	}
}
```
