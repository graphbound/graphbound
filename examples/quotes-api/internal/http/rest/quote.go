package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/domain"
	"go.uber.org/zap"
)

type QuoteController[T any] interface {
	GetQuote(c *T)
}

type quoteController struct {
	getQuoteUseCase domain.GetQuoteUseCase
	logger          *zap.SugaredLogger
}

var _ (QuoteController[gin.Context]) = (*quoteController)(nil)

func ProvideQuoteController(
	getQuoteUseCase domain.GetQuoteUseCase,
	logger *zap.SugaredLogger,
) *quoteController {
	return &quoteController{
		getQuoteUseCase: getQuoteUseCase,
		logger:          logger,
	}
}

func (qc quoteController) GetQuote(c *gin.Context) {
	resp, err := qc.getQuoteUseCase.Execute(c.Request.Context())
	if err != nil {
		qc.logger.Errorw("error getting quote",
			"error", err,
		)
	}
	qc.logger.Infof("quote: \"%s\"", resp.Value)
	c.JSON(200, resp)
}
