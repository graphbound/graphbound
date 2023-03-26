package requestid

import (
	ginrequestid "github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func NewServerPlugin() gin.HandlerFunc {
	return ginrequestid.New(ginrequestid.WithHandler(
		func(c *gin.Context, rid string) {
			c.Request = c.Request.WithContext(
				NewContext(c.Request.Context(), rid),
			)
		},
	))
}
