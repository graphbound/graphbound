package main

func main() {
	/*
		defer func() {
			if err := serverTraceProvider.Shutdown(context.Background()); err != nil {
				logger.Errorf("Error shutting down api tracer provider: %w", err)
			}
			if err := yeAPITracerProvider.Shutdown(context.Background()); err != nil {
				logger.Errorf("Error shutting down datasource tracer provider: %w", err)
			}
			if err := logger.Sync(); err != nil {
				logger.Errorf("Error flushing logger: %w", err)
			}
		}()

		health.WithServer(api.engine,
			health.NewServerComponent("api", "1.0.0"),
			health.NewHTTPDSHealthCheck("ye-api", "https://api.kanye.rest"),
		)
		err = api.engine.Run()
		if err != nil {
			logger.Panicf("Error starting api: %w", err)
		}
	*/
	api, err := initializeAPI()
	if err != nil {
		panic(err)
	}
	api.server.GET("/", api.quoteController.GetQuote)
	if err := api.server.Run(); err != nil {
		panic(err)
	}
}
