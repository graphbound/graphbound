package quote

import (
	"github.com/google/wire"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/domain"
)

var (
	GetQuoteUseCaseProviderSet = wire.NewSet(
		ProvideGetQuoteUseCase,
		wire.Bind(
			new(domain.GetQuoteUseCase),
			new(*getQuoteUseCase),
		),
	)
)
