package log

import (
	"time"

	"github.com/graphbound/graphbound/pkg/httpds"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewHTTPDSPlugin creates a logging plugin for HTTP data sources. Logs
// the HTTP request, response and error, if any.
func NewHTTPDSPlugin(logger *zap.Logger) httpds.Plugin {
	return func(c *httpds.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		fields := []zapcore.Field{
			zap.Int("status", c.Response.StatusCode),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Duration("latency", latency),
		}

		rid := c.Request.Header.Get("X-Request-ID")
		if rid != "" {
			fields = append(fields, zap.String("requestID", rid))
		}

		if c.Error != nil {
			logger.Error(c.Error.Error(), fields...)
		} else {
			logger.Info(path, fields...)
		}
	}
}
