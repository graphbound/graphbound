package log

import (
	"time"

	"github.com/graphbound/graphbound/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new logger
func NewLogger(env config.AppEnvironment) *zap.SugaredLogger {
	var logger *zap.Logger
	if config.IsProduction(env) {
		logger = zap.Must(zap.NewProduction())
	} else {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.Kitchen)
		logger = zap.Must(cfg.Build())
	}
	return logger.Sugar()
}
