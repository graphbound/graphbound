package log

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graphbound/graphbound/pkg/ginctx"
	"go.uber.org/zap"
)

type handlerExtension struct {
	logger                     *zap.Logger
	introspectionOperationName string
}

var _ interface {
	graphql.ResponseInterceptor
	graphql.HandlerExtension
} = (*handlerExtension)(nil)

func (e handlerExtension) ExtensionName() string {
	return "Logging"
}

func (e handlerExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (e handlerExtension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
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

	if rctx.Operation != nil {
		op := string(rctx.Operation.Operation)
		fields = append(fields, zap.String("type", op))
		e.logger.Info(op, fields...)
	} else {
		for _, err := range resp.Errors {
			msg := fmt.Sprintf("%s: %s", err.Rule, err.Message)
			e.logger.Error(msg, fields...)
		}
	}

	return resp
}
