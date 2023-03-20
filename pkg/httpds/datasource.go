package httpds

import (
	"context"
	"net/http"
)

type Plugin func(c *Context)

type DataSource struct {
	client  *http.Client
	plugins []Plugin
}

func New(client *http.Client, plugins ...Plugin) DataSource {
	return DataSource{
		client:  client,
		plugins: plugins,
	}
}

func (ds *DataSource) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	dsCtx := NewContext(ds.plugins...)
	dsCtx.Use(ds.do(ctx, req))
	dsCtx.Handle()
	return dsCtx.Response, dsCtx.Error
}

func (ds *DataSource) do(ctx context.Context, req *http.Request) func(c *Context) {
	return func(c *Context) {
		c.Ctx = ctx
		c.Request = req
		c.Next()
		res, err := ds.client.Do(c.Request)
		if err != nil {
			c.Error = err
		}
		c.Response = res
	}
}
