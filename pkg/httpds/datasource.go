package httpds

import "net/http"

type Plugin interface {
	OnRequest(req *http.Request)
	OnResponse(req *http.Request, res *http.Response, err error)
}

type DataSource struct {
	client  *http.Client
	plugins []Plugin
}

func New(plugins ...Plugin) DataSource {
	return DataSource{
		client:  http.DefaultClient,
		plugins: plugins,
	}
}

func (ds DataSource) Do(req *http.Request) (*http.Response, error) {
	for _, p := range ds.plugins {
		p.OnRequest(req)
	}
	res, err := ds.client.Do(req)
	for _, p := range ds.plugins {
		p.OnResponse(req, res, err)
	}
	return res, err
}
