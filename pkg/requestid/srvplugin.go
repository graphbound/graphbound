package requestid

import (
	ginrequestid "github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// NewServerPlugin creates a new plugin that adds a request ID to the response
// using the X-Request-ID header. Passes the X-Request-ID value back to the
// caller if it's sent in the request headers.
func NewServerPlugin() gin.HandlerFunc {
	return ginrequestid.New(ginrequestid.WithHandler(
		func(c *gin.Context, rid string) {
			c.Request = c.Request.WithContext(
				NewContext(c.Request.Context(), rid),
			)
		},
	))
}
