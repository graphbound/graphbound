package shutdown

import (
	"context"
	"net/http"
	"syscall"
	"time"

	"github.com/skovtunenko/graterm"
	"go.uber.org/zap"
)

const (
	httpServerTerminationOrder graterm.Order = 0
	httpServerHookName         string        = "HTTPServer"
	httpServerTimeout          time.Duration = 5 * time.Second
	defaultWaitTimeout         time.Duration = 10 * time.Second
)

type Option func(*Shutdown)

type Shutdown struct {
	terminator *graterm.Terminator
	ctx        context.Context
	timeout    time.Duration
	logger     *zap.SugaredLogger
}

// Wait waits for all shutdown hooks to finish or times out
func (s *Shutdown) Wait() {
	if err := s.terminator.Wait(s.ctx, s.timeout); err != nil {
		s.logger.Errorf("graceful termination has timed out: %s", err)
	}
}

// New creates a graceful shutdown
func New(l *zap.SugaredLogger, opts ...Option) *Shutdown {
	s := &Shutdown{}

	terminator, ctx := graterm.NewWithSignals(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	s.terminator = terminator
	s.ctx = ctx
	s.timeout = defaultWaitTimeout
	s.logger = l
	s.terminator.SetLogger(&logger{s.logger})

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithTimeout configures the graceful shutdown timeout
func WithTimeout(timeout time.Duration) Option {
	return func(s *Shutdown) {
		s.timeout = timeout
	}
}

// WithTimeout adds a new hook to the graceful stutdown. This hook waits for
// an HTTP server to shutdown.
func WithServer(srv *http.Server) Option {
	return func(s *Shutdown) {
		s.terminator.
			WithOrder(httpServerTerminationOrder).
			WithName(httpServerHookName).
			Register(httpServerTimeout, func(ctx context.Context) {
				if err := srv.Shutdown(ctx); err != nil {
					s.logger.Errorf("server forced to shutdown: %s", err)
				}
			})
	}
}
