package main

import (
  "net/http"
  "fmt"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

// LVM collector, listen to port 9080 path /metrics
func main() {

  lvmCollector := newLvmCollector()
  prometheus.MustRegister(lvmCollector)

  http.Handle("/metrics", promhttp.Handler())
  fmt.Println("Beginning to serve on port :9080")
  fmt.Println(http.ListenAndServe(":9080", nil))
}
