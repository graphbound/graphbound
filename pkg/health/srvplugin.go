package health

import (
	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v5"
)

const (
	readyzPath = "/readyz"
	livezPath  = "/livez"
)

// WithServer adds a live check endpoint and ready check endpoint to a given
// server
func WithServer(srv *gin.Engine, component health.Component, checks ...health.Config) {
	srv.GET(livezPath, newLivezHandler(component))
	srv.GET(readyzPath, newReadyzHandler(component, checks...))
}

// newLivezHandler creates a live check. This check is used to check if the
// server is available and alive
func newLivezHandler(component health.Component) gin.HandlerFunc {
	h, err := health.New(
		health.WithComponent(component),
		health.WithSystemInfo(),
	)
	if err != nil {
		panic(err)
	}

	return gin.WrapH(h.Handler())
}

// newReadyzHandler creates a ready check. This check is used to check if the
// server is ready to use and serve the traffic.
func newReadyzHandler(component health.Component, checks ...health.Config) gin.HandlerFunc {
	h, err := health.New(
		health.WithComponent(component),
		health.WithSystemInfo(),
		health.WithChecks(checks...),
	)
	if err != nil {
		panic(err)
	}

	return gin.WrapH(h.Handler())
}
