package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/TirushV/tempDB/db"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)
	httpStatusCodes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_status_codes_total",
			Help: "Total count of HTTP status codes.",
		},
		[]string{"endpoint", "status_code"},
	)
	totalKeys = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "total_keys_in_db",
			Help: "Total number of keys in the DB.",
		},
	)
)

func init() {
	prometheus.MustRegister(requestLatency)
	prometheus.MustRegister(httpStatusCodes)
	prometheus.MustRegister(totalKeys)
}

func recordLatency(start time.Time, endpoint string) {
	duration := time.Since(start).Seconds()
	requestLatency.WithLabelValues(endpoint).Observe(duration)
}

func recordStatusCode(endpoint string, statusCode int) {
	httpStatusCodes.WithLabelValues(endpoint, fmt.Sprintf("%d", statusCode)).Inc()
}

func SetKeyHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	defer func() {
		recordLatency(startTime, "SetKeyHandler")
		recordStatusCode("SetKeyHandler", http.StatusOK)
	}()

	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		recordStatusCode("SetKeyHandler", http.StatusBadRequest)
		return
	}

	key, value := requestData["key"], requestData["value"]
	if key == "" || value == "" {
		http.Error(w, "Key and value are required", http.StatusBadRequest)
		recordStatusCode("SetKeyHandler", http.StatusBadRequest)
		return
	}

	db.SetKeyValue(key, value) // Use db package functions
	fmt.Fprintf(w, "Key-value pair set successfully: %s=%s", key, value)
}

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	defer func() {
		recordLatency(startTime, "GetKeyHandler")
	}()

	key := strings.TrimPrefix(r.URL.Path, "/get/")
	value, exists := db.GetKeyValue(key) // Use db package functions
	if exists {
		w.Write([]byte(value)) // Write the value as a plain string response
		recordStatusCode("GetKeyHandler", http.StatusOK)
	} else {
		http.NotFound(w, r)
		recordStatusCode("GetKeyHandler", http.StatusNotFound)
	}
}

func SearchKeysHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	defer func() {
		recordLatency(startTime, "SearchKeysHandler")
	}()

	prefix := r.URL.Query().Get("prefix")
	suffix := r.URL.Query().Get("suffix")

	// Use db package functions
	filteredKeys := db.SearchKeysByPrefixSuffix(prefix, suffix)

	// Sort the keys alphabetically
	sort.Strings(filteredKeys)

	response := map[string][]string{"keys": filteredKeys}
	json.NewEncoder(w).Encode(response)
}

func UpdateTotalKeysMetricPeriodically() {
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            total := float64(db.GetTotalKeys())
            totalKeys.Set(total)
        }
    }
}