package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "api_calls_total",
	Help: "The total number of processed API calls",
})

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		counter.Inc() // skip favicon.ico and friends
	}
	jsonOutput, _ := json.Marshal(map[string]string{"hello": "world"})
	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler()) // exposes the metrics Prometheus scrapes
	fmt.Println("Server :4444")
	http.ListenAndServe(":4444", nil)
}
