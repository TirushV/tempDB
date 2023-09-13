package main

import (
	"fmt"
	"net/http"

	"github.com/TirushV/tempDB/server"
)

func main() {
	http.HandleFunc("/get/", server.GetKeyHandler)
	http.HandleFunc("/key/set", server.SetKeyHandler)
	http.HandleFunc("/search", server.SearchKeysHandler)

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
