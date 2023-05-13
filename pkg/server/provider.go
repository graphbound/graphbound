package server

import (
	"context"
	"log"

	"github.com/google/wire"
	"github.com/graphbound/graphbound/pkg/config"
	"github.com/graphbound/graphbound/pkg/trace"
)

func ProvideTracerProvider(appEnvironment config.AppEnvironment) (*tracerProvider, func(), error) {
	tp := trace.NewTracerProvider(
		serviceName,
		appEnvironment,
	)
	cleanup := func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Println(err)
		}
	}
	return (*tracerProvider)(tp), cleanup, nil
}

var (
	RESTServerProviderSet = wire.NewSet(
		ProvideTracerProvider,
		NewRESTServer,
	)
	GraphQLServerProviderSet = wire.NewSet(
		ProvideTracerProvider,
		NewGraphQLServer,
	)
)
