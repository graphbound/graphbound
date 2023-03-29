package main

import (
	"context"

	"github.com/graphbound/graphbound/pkg/health"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/metric"
	"github.com/graphbound/graphbound/pkg/requestid"
	"github.com/graphbound/graphbound/pkg/trace"
)

func main() {
	logger := log.NewLogger(false)

	serverTraceProvider, err := trace.NewTracerProvider("server", false)
	if err != nil {
		logger.Panicf("Error creating tracer provider: %w", err)
	}

	yeAPITracerProvider, err := trace.NewTracerProvider("ye-api", false)
	if err != nil {
		logger.Panicf("Error creating tracer provider: %w", err)
	}

	defer func() {
		if err := serverTraceProvider.Shutdown(context.Background()); err != nil {
			logger.Errorf("Error shutting down server tracer provider: %w", err)
		}
		if err := yeAPITracerProvider.Shutdown(context.Background()); err != nil {
			logger.Errorf("Error shutting down datasource tracer provider: %w", err)
		}
		if err := logger.Sync(); err != nil {
			logger.Errorf("Error flushing logger: %w", err)
		}
	}()

	ye := NewYeAPI("https://api.kanye.rest",
		requestid.NewHTTPDSPlugin(),
		log.NewHTTPDSPlugin(logger.Desugar().Named("datasource")),
		trace.NewHTTPDSPlugin(yeAPITracerProvider),
		metric.NewHTTPDSPlugin(),
	)

	controller := NewQuoteController(ye, logger.Named("controller"))

	server := NewServer(controller)
	server.engine.Use(requestid.NewServerPlugin())
	server.engine.Use(log.NewServerPlugin(logger.Named("server")))
	server.engine.Use(trace.NewServerPlugin("server", serverTraceProvider)...)
	health.WithServer(server.engine,
		health.NewServerComponent("server", "1.0.0"),
		health.NewHTTPDSHealthCheck("ye-api", "https://api.kanye.rest"),
	)
	server.engine.Use(metric.NewServerPlugin())
	metric.WithServer(server.engine)
	server.engine.GET("/", controller.GetQuote)

	err = server.engine.Run()
	if err != nil {
		logger.Panicf("Error starting server: %w", err)
	}
}
