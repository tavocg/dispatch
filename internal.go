package dispatch

import "context"

type ctxKey int

const metadataKey ctxKey = iota

type commandMetadata struct {
	Privileged  bool
	Interactive bool
}

func withCtx(ctx context.Context, metadata commandMetadata) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, metadataKey, metadata)
}

func getMetadata(ctx context.Context) commandMetadata {
	if ctx == nil {
		return commandMetadata{}
	}

	metadata, ok := ctx.Value(metadataKey).(commandMetadata)
	if !ok {
		return commandMetadata{}
	}

	return metadata
}

func withPrivileged(ctx context.Context) context.Context {
	metadata := getMetadata(ctx)
	metadata.Privileged = true
	return withCtx(ctx, metadata)
}

func withInteractive(ctx context.Context) context.Context {
	metadata := getMetadata(ctx)
	metadata.Interactive = true
	return withCtx(ctx, metadata)
}

func isPrivileged(ctx context.Context) bool {
	return getMetadata(ctx).Privileged
}

func isInteractive(ctx context.Context) bool {
	return getMetadata(ctx).Interactive
}
