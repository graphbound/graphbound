package quote

import (
	"context"
	"fmt"

	"github.com/graphbound/graphbound/examples/quotes-api/internal/domain"
	"github.com/graphbound/graphbound/examples/quotes-api/pkg/yeapi"
	"go.uber.org/zap"
)

type getQuoteUseCase struct {
	yeAPI  yeapi.Client
	logger *zap.SugaredLogger
}

var _ (domain.GetQuoteUseCase) = (*getQuoteUseCase)(nil)

func ProvideGetQuoteUseCase(
	yeAPI yeapi.Client,
	logger *zap.SugaredLogger,
) *getQuoteUseCase {
	return &getQuoteUseCase{
		yeAPI:  yeAPI,
		logger: logger,
	}
}

func (u getQuoteUseCase) Execute(ctx context.Context) (*domain.Quote, error) {
	q, err := u.yeAPI.GetQuote(ctx)
	if err != nil {
		return nil, fmt.Errorf("Execute: can not get ye quote: %w", err)
	}
	return &domain.Quote{Value: q.Quote}, nil
}
