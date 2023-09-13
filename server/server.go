package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TirushV/tempDB/db"
)

func SetKeyHandler(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	key, value := requestData["key"], requestData["value"]
	if key == "" || value == "" {
		http.Error(w, "Key and value are required", http.StatusBadRequest)
		return
	}

	db.SetKeyValue(key, value) // Use db package functions
	fmt.Fprintf(w, "Key-value pair set successfully: %s=%s", key, value)
}

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/get/")
	value, exists := db.GetKeyValue(key) // Use db package functions
	if exists {
		w.Write([]byte(value)) // Write the value as a plain string response
	} else {
		http.NotFound(w, r)
	}
}

func SearchKeysHandler(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("prefix")
	suffix := r.URL.Query().Get("suffix")

	// Use db package functions
	filteredKeys := db.SearchKeysByPrefixSuffix(prefix, suffix)

	response := map[string][]string{"keys": filteredKeys}
	json.NewEncoder(w).Encode(response)
}
