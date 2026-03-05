package dispatch

import "context"

type ctxKey int

const interactiveKey ctxKey = iota

func withInteractive(ctx context.Context, interactive bool) context.Context {
	return context.WithValue(ctx, interactiveKey, interactive)
}

func isInteractive(ctx context.Context) bool {
	v, ok := ctx.Value(interactiveKey).(bool)
	return ok && v
}
