package ginctx

import "github.com/gin-gonic/gin"

func NewServerPlugin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(
			NewContext(c.Request.Context(), c),
		)
		c.Next()
	}
}
