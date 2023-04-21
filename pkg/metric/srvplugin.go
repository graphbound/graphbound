package metric

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	metricPath = "/metrics"
)

// WithServer adds a metrics endpoint to a given server
func WithServer(srv *gin.Engine) {
	srv.GET(metricPath, gin.WrapH(promhttp.Handler()))
}

// NewServerPlugin creates a metrics middleware for HTTP servers. Registers
// metrics for HTTP requests and responses.
func NewServerPlugin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == metricPath {
			c.Next()
			return
		}

		start := time.Now()
		reqHost := c.Request.URL.Host
		reqPath := c.FullPath()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)

		reqDur.WithLabelValues(c.Request.Method, reqHost, reqPath, status).Observe(elapsed)
		reqCnt.WithLabelValues(c.Request.Method, reqHost, reqPath, status).Inc()
		reqSz.Observe(float64(computeRequestSize(c.Request)))
		resSz.Observe(float64(c.Writer.Size()))
	}
}
