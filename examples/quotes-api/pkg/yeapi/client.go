package yeapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/graphbound/graphbound/pkg/httpds"
)

type ClientURL string

type GetQuoteResponse struct {
	Quote string `json:"quote"`
}

type Client interface {
	GetQuote(ctx context.Context) (*GetQuoteResponse, error)
}

type client struct {
	baseURL *url.URL
	httpds.DataSource
}

var _ (Client) = (*client)(nil)

func ProvideClient(rawURL ClientURL, plugins ...httpds.Plugin) *client {
	baseURL, err := url.Parse(string(rawURL))
	if err != nil {
		panic(err)
	}

	return &client{
		baseURL:    baseURL,
		DataSource: httpds.New(http.DefaultClient, plugins...),
	}
}

func (ds client) GetQuote(ctx context.Context) (*GetQuoteResponse, error) {
	rel := &url.URL{
		Path: "/",
	}
	url := ds.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := ds.DataSource.Do(ctx, "/", req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var jres *GetQuoteResponse
	err = json.NewDecoder(res.Body).Decode(&jres)
	return jres, err
}
