package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_processed_ops_total",
		Help: "The total number of processed events",
	})

	responseRate = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "responses",
		Name:      "http_response_rate",
		Help:      "The response rate of requests",
		Buckets:   prometheus.LinearBuckets(0, 25, 5),
	}, []string{"method", "path", "code"})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
