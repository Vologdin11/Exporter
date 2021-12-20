package exporter

import (
	"log"
	"net/http"
	"stp-exporter/internal/exporter"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run() {
	port := ":8081"
	r := prometheus.NewRegistry()
	for {
		collector, err := exporter.NewCollector()
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		r.MustRegister(collector)
		http.HandleFunc("/", root)
		http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
		log.Printf("HTTP Server is running on %s", port)
		log.Fatal(http.ListenAndServe(port, nil))
	}
}
