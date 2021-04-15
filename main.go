package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	recordMetrics()

	fmt.Printf("Starting server on port 2112\n")
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", home)
	http.ListenAndServe(":2112", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success":true}`))

	duration := time.Since(startTime)
	responseRate.WithLabelValues(r.Method, r.URL.Path, fmt.Sprintf("%d", http.StatusOK)).Observe(float64(duration.Microseconds()))
}
