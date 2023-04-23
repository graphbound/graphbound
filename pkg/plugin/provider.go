package plugin

import (
	"github.com/graphbound/graphbound/pkg/httpds"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/metric"
	"github.com/graphbound/graphbound/pkg/requestid"
	"github.com/graphbound/graphbound/pkg/trace"
	"go.uber.org/zap"
)

func ProvideHTTPDSPlugins(
	logger *zap.SugaredLogger,
	tracerProvider trace.TracerProvider,
) []httpds.Plugin {
	return []httpds.Plugin{
		requestid.NewHTTPDSPlugin(),
		log.NewHTTPDSPlugin(logger.Desugar()),
		trace.NewHTTPDSPlugin(tracerProvider),
		metric.NewHTTPDSPlugin(),
	}
}
