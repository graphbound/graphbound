package httpds

import (
	"context"
	"net/http"
)

type (
	// Plugin defines the handler used by datasource as return value.
	Plugin func(c *Context)

	// DataSource is a component that one can use to encapsulate fetching
	// data from a source, in this case: a REST API.
	// Create an instance of DataSource by using New() or Default().
	DataSource struct {
		client  *http.Client
		plugins []Plugin
	}
)

// New returns a new blank DataSource configured with a given HTTP client and
// a list of plugins.
func New(client *http.Client, plugins ...Plugin) DataSource {
	return DataSource{
		client:  client,
		plugins: plugins,
	}
}

// Do builds the request context, prepares the Context handler chain and runs
// the main handler. It provides the same exact features as [Client.Do].
//
// [Client.Do]: https://pkg.go.dev/net/http#Client.Do
func (ds *DataSource) Do(ctx context.Context, path string, req *http.Request) (*http.Response, error) {
	dsCtx := NewContext(ds.plugins...)
	dsCtx.fullPath = path
	dsCtx.Request = req.WithContext(ctx)
	dsCtx.Use(ds.do())
	dsCtx.Handle()
	return dsCtx.Response, dsCtx.Error
}

// do is the main handler for our data source Context. It runs the HTTP request,
// and saves the response and/or error in the current context.
func (ds *DataSource) do() func(c *Context) {
	return func(c *Context) {
		res, err := ds.client.Do(c.Request)
		if err != nil {
			c.Error = err
		}
		c.Response = res
		c.Next()
	}
}
