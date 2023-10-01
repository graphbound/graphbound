package shutdown

import (
	"github.com/skovtunenko/graterm"
	"go.uber.org/zap"
)

type logger struct {
	*zap.SugaredLogger
}

var _ (graterm.Logger) = (*logger)(nil)

// Printf implements graterm.Logger
func (l *logger) Printf(format string, v ...interface{}) {
	l.SugaredLogger.Debugf(format, v)
}
