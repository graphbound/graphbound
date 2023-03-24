package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new logger
func NewLogger(isProduction bool) *zap.SugaredLogger {
	var logger *zap.Logger
	if isProduction {
		logger = zap.Must(zap.NewProduction())
	} else {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.Kitchen)
		logger = zap.Must(cfg.Build())
	}
	return logger.Sugar()
}
