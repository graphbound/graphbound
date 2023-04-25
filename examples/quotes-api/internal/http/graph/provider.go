package graph

import (
	"github.com/google/wire"
	"github.com/graphbound/graphbound/examples/quotes-api/internal/http/graph/generated"
)

var (
	ResolverProviderSet = wire.NewSet(
		ProvideResolver,
		generated.NewExecutableSchema,
		wire.Struct(new(generated.Config), "Resolvers"),
		wire.Bind(new(generated.ResolverRoot), new(*Resolver)),
	)
)
