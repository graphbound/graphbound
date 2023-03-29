package log

import (
	"time"

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
