// main.go
package main

import (
    "fmt"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    retransmissions = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "tcp_retransmissions_total",
            Help: "Total number of TCP retransmissions",
        },
        []string{},
    )
)

func main() {
    prometheus.MustRegister(retransmissions)
    http.Handle("/metrics", promhttp.Handler())
    go func() {
        for {
            // Read the eBPF map (simplified; in practice, use BPF syscalls)
            retransmissions.WithLabelValues().Set(42) // Replace with actual value
            time.Sleep(5 * time.Second)
        }
    }()
    http.ListenAndServe(":8080", nil)
}

