package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const prefix = "reverse_proxy_service"

var (
	IncRequest = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "_hits",
			Help: "Total number of requests made to the proxy service",
		},
		[]string{"status", "path"},
	)

	SummeryTimes = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    prefix + "_times",
			Help:    "Response time in seconds for requests made to the proxy service",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"status", "method", "path"},
	)

	CreateMovieHits = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: prefix + "_movie_create",
			Help: "Show create movie hits",
		},
		[]string{"path"})

	ResponseSearchMovie = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "_movie_search",
			Help: "Show movie search results",
		},
		[]string{"service"})
)
