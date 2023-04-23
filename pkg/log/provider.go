package log

import (
	"log"

	"github.com/google/wire"
	"github.com/graphbound/graphbound/pkg/config"
	"go.uber.org/zap"
)

func ProvideLogger(env config.AppEnvironment) (*zap.SugaredLogger, func(), error) {
	logger := NewLogger(env)
	cleanup := func() {
		if err := logger.Sync(); err != nil {
			log.Println(err)
		}
	}
	return logger, cleanup, nil
}

var (
	LoggerProviderSet = wire.NewSet(
		ProvideLogger,
	)
)
