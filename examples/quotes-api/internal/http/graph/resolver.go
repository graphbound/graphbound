package graph

import (
	"github.com/graphbound/graphbound/examples/quotes-api/internal/domain"
)

type Resolver struct {
	getQuoteUseCase domain.GetQuoteUseCase
}

func ProvideResolver(
	getQuoteUseCase domain.GetQuoteUseCase,
) *Resolver {
	return &Resolver{
		getQuoteUseCase: getQuoteUseCase,
	}
}
