package main

import "github.com/gin-gonic/gin"

func main() {
	/*
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

		yeAPI := yeapi.ProvideClient("https://api.kanye.rest",
			requestid.NewHTTPDSPlugin(),
			log.NewHTTPDSPlugin(logger.Desugar().Named("YeAPI")),
			trace.NewHTTPDSPlugin(yeAPITracerProvider),
			metric.NewHTTPDSPlugin(),
		)

		getQuoteUseCase := quote.ProvideGetQuoteUseCase(yeAPI, logger.Named("GetQuoteUseCase"))

		quoteController := rest.ProvideQuoteController(getQuoteUseCase, logger.Named("QuoteController"))

		server := NewServer(quoteController)
		server.engine.Use(requestid.NewServerPlugin())
		server.engine.Use(log.NewServerPlugin(logger.Named("QuotesAPI")))
		server.engine.Use(trace.NewServerPlugin("server", serverTraceProvider)...)
		health.WithServer(server.engine,
			health.NewServerComponent("server", "1.0.0"),
			health.NewHTTPDSHealthCheck("ye-api", "https://api.kanye.rest"),
		)
		server.engine.Use(metric.NewServerPlugin())
		metric.WithServer(server.engine)
		server.engine.GET("/", quoteController.GetQuote)

		err = server.engine.Run()
		if err != nil {
			logger.Panicf("Error starting server: %w", err)
		}
	*/
	server, err := initializeServer()
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	engine.GET("/", server.quoteController.GetQuote)
	if err := engine.Run(); err != nil {
		panic(err)
	}
}
