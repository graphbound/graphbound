package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/graphbound/graphbound/pkg/config"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/metric"
	"github.com/graphbound/graphbound/pkg/requestid"
	"github.com/graphbound/graphbound/pkg/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

const serviceName = "Server"

type tracerProvider sdktrace.TracerProvider

func ProvideTracerProvider(appEnvironment config.AppEnvironment) *tracerProvider {
	tp := trace.NewTracerProvider(
		serviceName,
		appEnvironment,
	)
	return (*tracerProvider)(tp)
}

func ProvideServer(
	logger *zap.SugaredLogger,
	tracerProvider *tracerProvider,
) *gin.Engine {
	server := gin.New()
	server.Use(requestid.NewServerPlugin())
	server.Use(log.NewServerPlugin(logger.Named(serviceName)))
	server.Use(trace.NewServerPlugin(serviceName, (*sdktrace.TracerProvider)(tracerProvider))...)
	server.Use(metric.NewServerPlugin())

	metric.WithServer(server)

	return server
}

var (
	ServerProviderSet = wire.NewSet(
		ProvideTracerProvider,
		ProvideServer,
	)
)
