package requestid

import (
	"context"
)

type key string

const requestIDKey key = "X-Request-ID"

func FromContext(ctx context.Context) (string, bool) {
	rid, ok := ctx.Value(requestIDKey).(string)
	return rid, ok
}

func NewContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}
