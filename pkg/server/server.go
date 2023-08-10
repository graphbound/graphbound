package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/graphbound/graphbound/pkg/ginctx"
	"github.com/graphbound/graphbound/pkg/health"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/metric"
	"github.com/graphbound/graphbound/pkg/requestid"
	"github.com/graphbound/graphbound/pkg/trace"
	healthgo "github.com/hellofresh/health-go/v5"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

const (
	serviceName                string = "Server"
	introspectionOperationName string = "IntrospectionQuery"
)

type (
	tracerProvider sdktrace.TracerProvider

	Version string
)

func newServer(
	version Version,
	healthChecks []healthgo.Config,
) *gin.Engine {
	server := gin.New()
	server.Use(requestid.NewHTTPServerPlugin())

	metric.WithServer(server)
	health.WithServer(server,
		health.NewServerComponent(serviceName, string(version)),
		healthChecks...,
	)

	return server
}

func NewRESTServer(
	logger *zap.SugaredLogger,
	tracerProvider *tracerProvider,
	version Version,
	healthChecks []healthgo.Config,
) *gin.Engine {
	server := newServer(version, healthChecks)

	server.Use(log.NewRESTServerPlugin(logger.Named(serviceName)))
	server.Use(trace.NewRESTServerPlugin(serviceName, (*sdktrace.TracerProvider)(tracerProvider))...)
	server.Use(metric.NewRESTServerPlugin())

	return server
}

func NewGraphQLServer(
	logger *zap.SugaredLogger,
	tracerProvider *tracerProvider,
	version Version,
	healthChecks []healthgo.Config,
	es graphql.ExecutableSchema,
) *gin.Engine {
	server := newServer(version, healthChecks)
	gql := handler.NewDefaultServer(es)

	gql.Use(log.NewGraphQLServerPlugin(logger.Named(serviceName), introspectionOperationName))
	gql.Use(trace.NewGraphQLServerPlugin(serviceName, (*sdktrace.TracerProvider)(tracerProvider)))
	gql.Use(metric.NewGraphQLServerPlugin())

	graphqlHandler := gin.WrapH(gql)
	playgroundHandler := gin.WrapH(playground.Handler("GraphQL", "/query"))

	server.Use(ginctx.NewGraphQLServerPlugin())
	server.POST("/query", graphqlHandler)
	server.GET("/", playgroundHandler)

	return server
}
