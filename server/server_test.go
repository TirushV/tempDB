package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TirushV/tempDB/db"
)

func TestSetKeyHandler(t *testing.T) {
	// Create a test HTTP request with a JSON payload
	requestBody := `{"key": "test-key", "value": "test-value"}`
	req, err := http.NewRequest("POST", "/key/set", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the SetKeyHandler function to handle the request
	SetKeyHandler(rr, req)

	// Check the response status code and body
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := "Key-value pair set successfully: test-key=test-value"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}
}

func TestGetKeyHandler(t *testing.T) {
	// Initialize the database
	db.SetKeyValue("abc-1", "abc")

	// Create a test HTTP request to retrieve a key
	req, err := http.NewRequest("GET", "/get/abc-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the GetKeyHandler function to handle the request
	GetKeyHandler(rr, req)

	// Check the response status code and body
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := "abc"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}
}

func TestSearchKeysHandler(t *testing.T) {
	// Initialize the database with test data
	db.SetKeyValue("abc-1", "value1")
	db.SetKeyValue("abc-2", "value2")
	db.SetKeyValue("xyz-1", "value3")
	db.SetKeyValue("xyz-2", "value4")

	// Create a test HTTP request to search for keys by prefix
	req, err := http.NewRequest("GET", "/search?prefix=abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the SearchKeysHandler function to handle the request
	SearchKeysHandler(rr, req)

	// Check the response status code and body
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the response body for the expected keys
	expectedResponse := `{"keys":["abc-1","abc-2"]}`
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}
}
