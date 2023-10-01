//go:build wireinject

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/http/rest"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/quote"
	"github.com/graphbound/graphbound/examples/quotes-api/pkg/yeapi"
	"github.com/graphbound/graphbound/pkg/httpds"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/server"
	"github.com/hellofresh/health-go/v5"
	"go.uber.org/zap"
)

type API struct {
	quoteController rest.QuoteController[gin.Context]
	router          *gin.Engine
	logger          *zap.SugaredLogger
}

func ProvideAPI(
	quoteController rest.QuoteController[gin.Context],
	router *gin.Engine,
	logger *zap.SugaredLogger,
) (*API, error) {
	if logger == nil {
		return nil, fmt.Errorf("ProvideAPI: logger is nil")
	}

	return &API{
		quoteController: quoteController,
		router:          router,
		logger:          logger,
	}, nil
}

func ProvideHealthChecks(
	yeAPIHealthCheck yeapi.HealthCheck,
) []health.Config {
	return []health.Config{
		(health.Config)(yeAPIHealthCheck),
	}
}

func initializeAPI() (*API, func(), error) {
	panic(wire.Build(
		ProvideConfig,
		wire.FieldsOf(new(*Config),
			"AppEnvironment",
			"YeAPIURL",
		),
		log.LoggerProviderSet,
		yeapi.ClientProviderSet,
		quote.GetQuoteUseCaseProviderSet,
		rest.QuoteControllerProviderSet,
		server.RESTServerProviderSet,
		wire.Value(server.Version("1.0.0")),
		wire.Value([]httpds.Plugin(nil)),
		ProvideHealthChecks,
		ProvideAPI,
	))
}
