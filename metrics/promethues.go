package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

var pingGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "request_per_minutes",
		Help: "Number of request per_seconds",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	pingGauge.Inc()
	fmt.Fprintf(w, "pong")
}

func metrics() {
	prometheus.MustRegister(pingCounter, pingGauge)
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8081", nil)
}
