package requestid

import (
	"context"
	"net/http"
)

type key string

const (
	requestIDKey key = "X-Request-ID"
)

// FromHeader gets a request ID from a header.
func FromHeader(h http.Header) string {
	return h.Get(string(requestIDKey))
}

// FromContext gets a request ID from a context.
func FromContext(ctx context.Context) (string, bool) {
	rid, ok := ctx.Value(requestIDKey).(string)
	return rid, ok
}

// NewContext creates a new context with a request ID.
func NewContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}
