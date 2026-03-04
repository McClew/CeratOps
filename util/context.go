package util

import (
	"context"
)

type contextKey string

const (
	PathKey contextKey = "path"
)

// SetPath returns a new context with the path value set.
func SetPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, PathKey, path)
}

// GetPath retrieves the path value from the context.
func GetPath(ctx context.Context) string {
	if v := ctx.Value(PathKey); v != nil {
		if path, ok := v.(string); ok {
			return path
		}
	}
	return ""
}
