package trace

import (
	"context"

	"github.com/graphbound/graphbound/pkg/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const (
	requestIDKey = "http.request_id"
)

type TracerProvider oteltrace.TracerProvider

// NewTracerProvider creates a new tracer provider. One should one tracer
// provider per app/service (queues, HTTP data sources, databases, servers, etc.)
func NewTracerProvider(
	service string,
	appEnvironment config.AppEnvironment,
) *trace.TracerProvider {
	opts := []otlptracehttp.Option{}
	if config.IsProduction(appEnvironment) {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	client := otlptracehttp.NewClient(opts...)
	exporter, err := otlptrace.New(
		context.Background(),
		client,
	)
	if err != nil {
		panic(err)
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
		)),
	)
}
