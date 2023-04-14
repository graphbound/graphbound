package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	reqCnt = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of processed requests",
	}, []string{"method", "host", "path", "status"})
	reqDur = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of HTTP requests in seconds",
	}, []string{"method", "host", "path", "status"})
	reqSz = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "http_request_length_bytes",
		Help: "Content-Length of HTTP requests",
	})
	resSz = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "http_response_length_bytes",
		Help: "Content-Length of HTTP responses",
	})
)

func computeRequestSize(r *http.Request) int {
	s := 0
	s += len(r.URL.String())
	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
