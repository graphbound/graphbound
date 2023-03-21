package main

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/requestid"
)

func main() {
	logger := log.NewLogger(false)

	ye := NewYeAPI("https://api.kanye.rest",
		log.NewHTTPDSPlugin(logger.Desugar().Named("datasource")),
		requestid.NewHTTPDSPlugin(),
	)

	controller := NewQuoteController(ye, logger.Named("controller"))

	server := NewServer(controller, logger.Named("server"))
	server.engine.Use(ginzap.Ginzap(server.logger.Desugar(), time.RFC3339, true))
	server.engine.GET("/", controller.GetQuote)
	server.engine.Run()
}
