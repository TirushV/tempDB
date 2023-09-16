package db

import (
	"strings"
	"sync"
)

var (
	keyValueStore = make(map[string]string)
	mu            sync.Mutex
)

func SetKeyValue(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	keyValueStore[key] = value
}

func GetKeyValue(key string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	value, exists := keyValueStore[key]
	return value, exists
}

func SearchKeysByPrefixSuffix(prefix, suffix string) []string {
	mu.Lock()
	defer mu.Unlock()

	var filteredKeys []string
	for key := range keyValueStore {
		if (prefix == "" || strings.HasPrefix(key, prefix)) &&
			(suffix == "" || strings.HasSuffix(key, suffix)) {
			filteredKeys = append(filteredKeys, key)
		}
	}
	return filteredKeys
}

func GetTotalKeys() int {
	mu.Lock()
	defer mu.Unlock()
	return len(keyValueStore)
}