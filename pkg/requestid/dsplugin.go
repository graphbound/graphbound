package requestid

import (
	"github.com/graphbound/graphbound/pkg/httpds"
)

// NewHTTPDSPlugin creates a new plugin that injects request IDs into the HTTP
// request. Adds the X-Request-ID header if the request context has one.
func NewHTTPDSPlugin() httpds.Plugin {
	return func(c *httpds.Context) {
		if rid, ok := FromContext(c.Request.Context()); ok {
			c.Request.Header.Add("X-Request-ID", rid)
		}
		c.Next()
	}
}
