package analytics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Define a simple counter metric
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "event_total",
			Help: "Total Number of Event",
		},
		[]string{"event", "meta_data"},
	)

	// requestHist = prometheus.NewHistogramVec(opts prometheus.HistogramOpts, labelNames []string)
	Registry = prometheus.NewRegistry()
)

// IncrementEventCount increments the event count metric
func IncrementEventCount(event, metaData string) {
	mapCount := map[string]string{
		"event":     event,
		"meta_data": metaData,
	}
	requestCount.With(mapCount).Inc()
}

func init() {
	// Register the metric with Prometheus
	Registry.MustRegister(requestCount)
}
