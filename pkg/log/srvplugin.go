package log

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewServerPlugin creates a logging middleware for HTTP servers. Logs
// the HTTP request, response and error, if any.
func NewServerPlugin(logger *zap.SugaredLogger) gin.HandlerFunc {
	return ginzap.Ginzap(
		logger.Desugar(),
		time.RFC3339,
		true,
	)
}

// NewGraphQLServerPlugin creates a logging middleware for GraphQL servers. Logs
// GraphQL operations, responses and errors, if any.
func NewGraphQLServerPlugin(
	logger *zap.SugaredLogger,
	introspectionOperationName string,
) graphql.HandlerExtension {
	return handlerExtension{
		logger:                     logger.Desugar(),
		introspectionOperationName: introspectionOperationName,
	}
}
