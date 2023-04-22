package yeapi

import (
	"github.com/google/wire"
)

var (
	ClientProvider = wire.NewSet(
		ProvideClient,
		wire.Bind(
			new(Client),
			new(*client),
		),
	)
)
