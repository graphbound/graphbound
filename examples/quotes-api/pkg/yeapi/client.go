package yeapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/graphbound/graphbound/pkg/httpds"
	"github.com/graphbound/graphbound/pkg/plugin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

const serviceName = "YeAPI"

type (
	ClientURL string

	tracerProvider sdktrace.TracerProvider

	GetQuoteResponse struct {
		Quote string `json:"quote"`
	}

	Client interface {
		GetQuote(ctx context.Context) (*GetQuoteResponse, error)
	}

	client struct {
		baseURL *url.URL
		logger  *zap.SugaredLogger
		httpds.DataSource
	}
)

var _ (Client) = (*client)(nil)

func ProvideClient(
	rawURL ClientURL,
	logger *zap.SugaredLogger,
	tracerProvider *tracerProvider,
	plugins ...httpds.Plugin,
) *client {
	baseURL, err := url.Parse(string(rawURL))
	if err != nil {
		panic(err)
	}

	l := logger.Named(serviceName)
	p := plugin.ProvideHTTPDSPlugins(l, (*sdktrace.TracerProvider)(tracerProvider))
	p = append(p, plugins...)

	return &client{
		baseURL:    baseURL,
		logger:     l,
		DataSource: httpds.New(http.DefaultClient, p...),
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
