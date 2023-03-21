package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type quoteController struct {
	yeAPI  YeAPI
	logger *zap.SugaredLogger
}

func NewQuoteController(
	yeAPI YeAPI,
	logger *zap.SugaredLogger,
) *quoteController {
	return &quoteController{
		yeAPI:  yeAPI,
		logger: logger,
	}
}

func (qc quoteController) GetQuote(c *gin.Context) {
	resp, err := qc.yeAPI.GetQuote(context.Background())
	if err != nil {
		qc.logger.Errorw("error getting quote",
			"error", err,
		)
	}
	qc.logger.Infof("quote: \"%s\"", resp.Quote)
	c.JSON(200, resp)
}
