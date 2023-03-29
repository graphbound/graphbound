package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/graphbound/graphbound/pkg/requestid"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// NewServerPlugin creates a tracing middleware for HTTP servers. Traces
// the HTTP request and injects the tracer into the request context.
func NewServerPlugin(service string, provider trace.TracerProvider) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		otelgin.Middleware(
			service,
			otelgin.WithTracerProvider(provider),
		),
		func(c *gin.Context) {
			if rid := requestid.FromHeader(c.Request.Header); rid != "" {
				s := trace.SpanFromContext(c.Request.Context())
				s.SetAttributes(attribute.String(requestIDKey, rid))
			}
			c.Next()
		},
	}
}
