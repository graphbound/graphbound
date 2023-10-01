package main

import (
	"net/http"

	"github.com/graphbound/graphbound/pkg/shutdown"
)

func main() {
	api, cleanup, err := initializeAPI()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	api.router.GET("/", api.quoteController.GetQuote)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: api.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			api.logger.Fatalf("listen: %s", err)
		}
	}()

	shutdown.New(api.logger,
		shutdown.WithServer(srv),
	).Wait()
}
