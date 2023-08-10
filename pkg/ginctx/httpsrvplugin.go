package ginctx

import "github.com/gin-gonic/gin"

// NewGraphQLServerPlugin creates a context plugin for GraphQL servers. Injects
// gin's context into the HTTP request context.
func NewGraphQLServerPlugin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(
			NewContext(c.Request.Context(), c),
		)
		c.Next()
	}
}
