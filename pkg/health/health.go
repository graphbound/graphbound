package health

import (
	"github.com/hellofresh/health-go/v5"
	"github.com/hellofresh/health-go/v5/checks/http"
)

// NewServerComponent creates a health component for a server
func NewServerComponent(name, version string) health.Component {
	return health.Component{
		Name:    name,
		Version: version,
	}
}

// NewHTTPDSHealthCheck creates a health check config for an HTTP datasource
func NewHTTPDSHealthCheck(name, url string) health.Config {
	return health.Config{
		Name: name,
		Check: http.New(http.Config{
			URL: url,
		}),
	}
}
