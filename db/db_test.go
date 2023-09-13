package db

import (
	"testing"
)

func TestSetKeyValue(t *testing.T) {
	// Initialize the database
	keyValueStore = make(map[string]string)

	key := "test-key"
	value := "test-value"
	SetKeyValue(key, value)

	retrievedValue, exists := GetKeyValue(key)
	if !exists {
		t.Errorf("Expected key %s to exist in the database", key)
	}

	if retrievedValue != value {
		t.Errorf("Expected value %s, got %s", value, retrievedValue)
	}
}

func TestGetKeyValue(t *testing.T) {
	// Initialize the database
	keyValueStore = make(map[string]string)

	key := "test-key"
	value := "test-value"
	SetKeyValue(key, value)

	retrievedValue, exists := GetKeyValue(key)
	if !exists {
		t.Errorf("Expected key %s to exist in the database", key)
	}

	if retrievedValue != value {
		t.Errorf("Expected value %s, got %s", value, retrievedValue)
	}
}

func TestSearchKeysByPrefixSuffix(t *testing.T) {
	// Initialize the database with test data
	keyValueStore = map[string]string{
		"abc-1": "value1",
		"abc-2": "value2",
		"xyz-1": "value3",
		"xyz-2": "value4",
	}

	prefix := "abc"
	suffix := "-1"
	filteredKeys := SearchKeysByPrefixSuffix(prefix, suffix)

	expectedKeys := []string{"abc-1"}
	for i, key := range expectedKeys {
		if filteredKeys[i] != key {
			t.Errorf("Expected key %s, got %s", key, filteredKeys[i])
		}
	}
}
