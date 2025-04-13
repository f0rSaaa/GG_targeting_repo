package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Request metrics
	RequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "targeting_request_total",
			Help: "Total number of requests received",
		},
		[]string{"status"},
	)

	// Campaign metrics
	CampaignsReturned = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "targeting_campaigns_returned",
			Help:    "Number of campaigns returned per request",
			Buckets: []float64{0, 1, 2, 5, 10, 20, 50},
		},
	)

	// Latency metrics
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "targeting_request_duration_seconds",
			Help:    "Time taken to process request",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)

	// Database metrics
	DBQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "targeting_db_query_duration_seconds",
			Help:    "Time taken for database queries",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query_type"},
	)
)
