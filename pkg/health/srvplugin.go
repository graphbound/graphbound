package health

import (
	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v5"
)

const (
	readyzPath = "/readyz"
	livezPath  = "/livez"
)

func WithServer(srv *gin.Engine, component health.Component, checks ...health.Config) {
	srv.GET(livezPath, newLivezHandler(component))
	srv.GET(readyzPath, newReadyzHandler(component, checks...))
}

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
