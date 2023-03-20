package requestid

import (
	"github.com/graphbound/graphbound/pkg/httpds"
)

func NewHTTPDSPlugin() httpds.Plugin {
	return func(c *httpds.Context) {
		if rid, ok := FromContext(c.Ctx); ok {
			c.Request.Header.Add("X-Request-ID", rid)
		}
		c.Next()
	}
}
