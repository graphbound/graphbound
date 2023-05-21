package log

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graphbound/graphbound/pkg/ginctx"
	"go.uber.org/zap"
)

type loggerHandlerExtension struct {
	logger                     *zap.Logger
	introspectionOperationName string
}

var _ interface {
	graphql.ResponseInterceptor
	graphql.HandlerExtension
} = (*loggerHandlerExtension)(nil)

func (e loggerHandlerExtension) ExtensionName() string {
	return "Logging"
}

func (e loggerHandlerExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (e loggerHandlerExtension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	start := time.Now()

	resp := next(ctx)

	end := time.Now()
	latency := end.Sub(start)

	rctx := graphql.GetOperationContext(ctx)
	if rctx.OperationName == e.introspectionOperationName {
		return resp
	}

	fields := []zap.Field{
		zap.String("query", rctx.RawQuery),
		zap.Any("variables", rctx.Variables),
		zap.Duration("latency", latency),
	}

	op := rctx.OperationName
	if rctx.Operation != nil {
		t := string(rctx.Operation.Operation)
		op += ":" + t
		fields = append(fields, zap.String("type", t))
	}

	c, ok := ginctx.FromContext(ctx)
	if ok {
		fields = append(fields,
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		)

		rid := c.GetHeader("X-Request-ID")
		if rid != "" {
			fields = append(fields, zap.String("requestID", rid))
		}
	}

	e.logger.Info(op, fields...)

	return resp
}
