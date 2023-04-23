//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/http/rest"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/quote"
	"github.com/graphbound/graphbound/examples/quotes-api/pkg/yeapi"
	"github.com/graphbound/graphbound/pkg/httpds"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/server"
	"github.com/hellofresh/health-go/v5"
)

type API struct {
	quoteController rest.QuoteController[gin.Context]
	server          *gin.Engine
}

func ProvideAPI(
	quoteController rest.QuoteController[gin.Context],
	server *gin.Engine,
) *API {
	return &API{
		quoteController: quoteController,
		server:          server,
	}
}

func ProvideHealthChecks(
	yeAPIHealthCheck yeapi.HealthCheck,
) []health.Config {
	return []health.Config{
		(health.Config)(yeAPIHealthCheck),
	}
}

func initializeAPI() (*API, error) {
	wire.Build(
		ProvideConfig,
		wire.FieldsOf(new(*Config),
			"AppEnvironment",
			"YeAPIURL",
		),
		log.NewLogger,
		yeapi.ClientProviderSet,
		quote.GetQuoteUseCaseProviderSet,
		rest.QuoteControllerProviderSet,
		server.ServerProviderSet,
		wire.Value(server.Version("1.0.0")),
		wire.Value([]httpds.Plugin(nil)),
		ProvideHealthChecks,
		ProvideAPI,
	)
	return &API{}, nil
}
