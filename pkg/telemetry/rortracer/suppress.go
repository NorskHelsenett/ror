package rortracer

import "context"

type suppressKey struct{}

// SuppressTracing returns a new context that prevents StartSpan from creating
// real spans. Any call to StartSpan with this context (or a child of it) will
// return a noop span. This is useful for suppressing traces on health-check
// endpoints, internal metrics scrapes, or other high-frequency low-value paths.
func SuppressTracing(ctx context.Context) context.Context {
	return context.WithValue(ctx, suppressKey{}, true)
}

// IsTracingSuppressed reports whether the context has been marked to suppress
// tracing via SuppressTracing.
func IsTracingSuppressed(ctx context.Context) bool {
	v, _ := ctx.Value(suppressKey{}).(bool)
	return v
}
