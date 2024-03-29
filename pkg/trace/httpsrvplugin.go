package trace

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/graphbound/graphbound/pkg/requestid"
	"github.com/ravilushqa/otelgqlgen"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// NewRESTServerPlugin creates a tracing plugin for REST servers. Traces the
// HTTP request and injects the tracer into the request context.
func NewRESTServerPlugin(service string, provider trace.TracerProvider) []gin.HandlerFunc {
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

// NewGraphQLServerPlugin creates a tracing plugin for GraphQL servers. Traces
// the GraphQL request and injects the tracer into the request context.
func NewGraphQLServerPlugin(service string, provider trace.TracerProvider) graphql.HandlerExtension {
	return otelgqlgen.Middleware(otelgqlgen.WithTracerProvider(provider))
}
