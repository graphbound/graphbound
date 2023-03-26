package log

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewServerPlugin(logger *zap.SugaredLogger) gin.HandlerFunc {
	return ginzap.Ginzap(
		logger.Desugar(),
		time.RFC3339,
		true,
	)
}
