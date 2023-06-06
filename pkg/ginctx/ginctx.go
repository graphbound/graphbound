package ginctx

import (
	"context"

	"github.com/gin-gonic/gin"
)

const srvContextKey key = "ServerContext"

type key string

// NewContext creates a new context with gin context.
func NewContext(ctx context.Context, c *gin.Context) context.Context {
	return context.WithValue(ctx, srvContextKey, c)
}

// FromContext gets a gin context from a context.
func FromContext(ctx context.Context) (*gin.Context, bool) {
	c, ok := ctx.Value(srvContextKey).(*gin.Context)
	return c, ok
}
