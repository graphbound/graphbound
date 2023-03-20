package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/graphbound/graphbound/pkg/httpds"
	"github.com/graphbound/graphbound/pkg/log"
	"github.com/graphbound/graphbound/pkg/requestid"
)

type GetQuoteResponse struct {
	Quote string `json:"quote"`
}

type YeAPI interface {
	GetQuote() (*GetQuoteResponse, error)
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

	res, err := ds.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var jres *GetQuoteResponse
	err = json.NewDecoder(res.Body).Decode(&jres)
	return jres, err
}

func main() {
	logger := log.NewLogger(false)

	ye := NewYeAPI("https://api.kanye.rest",
		log.NewHTTPDSPlugin(logger.Desugar().Named("YeAPI")),
		requestid.NewHTTPDSPlugin(),
	)

	ctx := requestid.NewContext(context.Background(), uuid.New().String())

	resp, err := ye.GetQuote(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Quote)
}
