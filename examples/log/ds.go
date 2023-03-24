package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/graphbound/graphbound/pkg/httpds"
)

type GetQuoteResponse struct {
	Quote string `json:"quote"`
}

type YeAPI interface {
	GetQuote(ctx context.Context) (*GetQuoteResponse, error)
}

type yeAPI struct {
	baseURL *url.URL
	httpds.DataSource
}

func NewYeAPI(rawURL string, plugins ...httpds.Plugin) *yeAPI {
	baseURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	return &yeAPI{
		baseURL:    baseURL,
		DataSource: httpds.New(http.DefaultClient, plugins...),
	}
}

func (ds yeAPI) GetQuote(ctx context.Context) (*GetQuoteResponse, error) {
	rel := &url.URL{
		Path: "/",
	}
	url := ds.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := ds.DataSource.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var jres *GetQuoteResponse
	err = json.NewDecoder(res.Body).Decode(&jres)
	return jres, err
}
