package trace

import (
	"github.com/graphbound/graphbound/pkg/httpds"
	"github.com/graphbound/graphbound/pkg/requestid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/semconv/v1.17.0/httpconv"
	"go.opentelemetry.io/otel/trace"
)

const (
	tracerName = "github.com/graphbound/graphbound/pkg/trace"
)

// NewHTTPDSPlugin creates a tracing middleware for HTTP data sources. Traces
// the HTTP request.
func NewHTTPDSPlugin(provider trace.TracerProvider) httpds.Plugin {
	tracer := provider.Tracer(
		tracerName,
	)

	return func(c *httpds.Context) {
		attrs := httpconv.ClientRequest(c.Request)
		if rid := requestid.FromHeader(c.Request.Header); rid != "" {
			attrs = append(attrs, attribute.String(requestIDKey, rid))
		}
		opts := []trace.SpanStartOption{
			trace.WithAttributes(attrs...),
			trace.WithSpanKind(trace.SpanKindClient),
		}
		spanName := c.Request.URL.Path
		ctx, span := tracer.Start(c.Request.Context(), spanName, opts...)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)

		c.Next()

		status := c.Response.StatusCode
		span.SetStatus(httpconv.ClientStatus(status))
		span.SetAttributes(
			httpconv.ClientResponse(c.Response)...,
		)
		if c.Error != nil {
			span.SetAttributes(
				attribute.String("data.source.error", c.Error.Error()),
			)
		}
	}
}
