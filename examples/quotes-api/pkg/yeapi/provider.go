package yeapi

import (
	"github.com/google/wire"
	"github.com/graphbound/graphbound/pkg/config"
	"github.com/graphbound/graphbound/pkg/trace"
)

func ProvideTracerProvider(appEnvironment config.AppEnvironment) *tracerProvider {
	tp := trace.NewTracerProvider(
		serviceName,
		appEnvironment,
	)

	return (*tracerProvider)(tp)
}

var (
	ClientProviderSet = wire.NewSet(
		ProvideTracerProvider,
		ProvideClientHealthCheck,
		ProvideClient,
		wire.Bind(
			new(Client),
			new(*client),
		),
	)
)
