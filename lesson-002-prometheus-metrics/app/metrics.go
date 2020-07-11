package app

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	responsesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "http",
			Name:      "responses_total",
			Help:      "Total http responses",
		},
		[]string{
			// HTTP status code.
			"code",
		},
	)
)

// RegisterMetrics registers all metrics for the server.
func RegisterMetrics() {
	prometheus.MustRegister(responsesTotal)
}
