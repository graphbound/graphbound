package httpds

import (
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

func (ds *DataSource) Do(req *http.Request) (*http.Response, error) {
	ctx := NewContext(ds.plugins...)
	ctx.Use(ds.do(req))
	ctx.Handle()
	return ctx.Response, ctx.Error
}

func (ds *DataSource) do(req *http.Request) func(c *Context) {
	return func(c *Context) {
		c.Request = req
		c.Next()
		res, err := ds.client.Do(c.Request)
		if err != nil {
			c.Error = err
		}
		c.Response = res
	}
}
