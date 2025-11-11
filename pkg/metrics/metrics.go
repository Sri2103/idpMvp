// pkg/metrics/metrics.go
// Used by: Every HTTP/gRPC service to emit metrics.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "idp_requests_total",
		Help: "Total requests processed by service",
	},
	[]string{"service", "endpoint", "status"},
)

func Init() {
	prometheus.MustRegister(RequestCount)
}
