package main

import (
	"fmt"
	"net/http"

	"github.com/TirushV/tempDB/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go server.UpdateTotalKeysMetricPeriodically()

	http.HandleFunc("/get/", server.GetKeyHandler)
	http.HandleFunc("/key/set", server.SetKeyHandler)
	http.HandleFunc("/search", server.SearchKeysHandler)
	http.Handle("/metrics", promhttp.Handler())

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}