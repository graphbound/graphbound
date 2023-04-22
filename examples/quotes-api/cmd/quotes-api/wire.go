//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/http/rest"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/quote"
	"github.com/graphbound/graphbound/examples/quotes-api/pkg/yeapi"
	"github.com/graphbound/graphbound/pkg/config"
	"github.com/graphbound/graphbound/pkg/httpds"
	"github.com/graphbound/graphbound/pkg/log"
)

type server struct {
	quoteController rest.QuoteController[gin.Context]
}

func ProvideServer(
	quoteController rest.QuoteController[gin.Context],
) *server {
	return &server{
		quoteController: quoteController,
	}
}

func initializeServer() (*server, error) {
	wire.Build(
		ProvideConfig,
		log.NewLogger,
		yeapi.ClientProvider,
		quote.GetQuoteUseCaseProvider,
		rest.QuoteControllerProvider,
		wire.FieldsOf(new(*Config), "YeAPIURL"),
		wire.Value([]httpds.Plugin(nil)),
		wire.Value(config.AppEnvironment("development")),
		ProvideServer,
	)
	return &server{}, nil
}
