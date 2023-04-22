package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var (
	QuoteControllerProvider = wire.NewSet(
		ProvideQuoteController,
		wire.Bind(
			new(QuoteController[gin.Context]),
			new(*quoteController),
		),
	)
)
