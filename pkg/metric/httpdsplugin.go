package metric

import (
	"strconv"
	"time"

	"github.com/graphbound/graphbound/pkg/httpds"
)

// NewHTTPDSPlugin creates a metrics plugin for HTTP data sources. Registers
// the HTTP request method, host, path and status.
func NewHTTPDSPlugin() httpds.Plugin {
	return func(c *httpds.Context) {
		start := time.Now()
		reqHost := c.Request.URL.Host
		reqPath := c.FullPath()

		c.Next()

		status := strconv.Itoa(c.Response.StatusCode)
		elapsed := float64(time.Since(start)) / float64(time.Second)

		reqDur.WithLabelValues(c.Request.Method, reqHost, reqPath, status).Observe(elapsed)
		reqCnt.WithLabelValues(c.Request.Method, reqHost, reqPath, status).Inc()
	}
}
