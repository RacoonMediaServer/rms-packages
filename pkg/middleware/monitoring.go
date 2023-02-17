package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PanicCounter                prometheus.Counter
	RequestsCounter             *prometheus.CounterVec
	UnauthorizedRequestsCounter prometheus.Counter
	RequestDuration             prometheus.Histogram
)

func init() {
	PanicCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "rms",
		Name:      "panics",
		Help:      "The total number of panic happens",
	})

	RequestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "rms",
		Name:      "server_requests",
		Help:      "The total number of processed requests",
	}, []string{"method", "code"})

	UnauthorizedRequestsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "rms",
		Name:      "server_unauthorized_requests",
		Help:      "The total number of unauthorized requests",
	})

	RequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "rms",
		Name:      "server_requests_duration",
		Help:      "Duration of requests",
	})
}
