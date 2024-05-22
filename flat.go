// Package goflat provides functions for flattening and unflattening JSON objects.
// It supports flattening JSON arrays as well as nested maps.
package goflat

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Options represents the options for flattening and unflattening JSON.
type Options struct {
	KeyDelimiter string // The delimiter to use for separating keys in the flattened structure
	MaxDepth     int    // The maximum depth for flattening
}

// DefaultOptions returns the default options for flattening and unflattening JSON.
func DefaultOptions() Options {
	return Options{
		KeyDelimiter: ".",
		MaxDepth:     -1, // -1 indicates no maximum depth
	}
}

// FlattenJSON flattens a JSON object into a map[string]interface{} using the specified options.
// It supports flattening JSON arrays as well.
//
// Example:
//
//	data := []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`)
//	options := DefaultOptions()
//	flattened, err := FlattenJSON(data, options)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	fmt.Println(flattened)
//
// Output:
//
//	map[address.city:New York address.state:NY age:30 name:John]
func FlattenJSON(data []byte, options Options) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	flattened := make(map[string]interface{})
	flatten("", result, flattened, options.KeyDelimiter, options.MaxDepth)
	return flattened, nil
}

// FlattenMap flattens a map[string]interface{} into a map[string]interface{} using the specified options.
// It supports flattening nested maps as well.
//
// Example:
//
//	data := map[string]interface{}{
//		"name": "John",
//		"age": 30,
//		"address": map[string]interface{}{
//			"city":  "New York",
//			"state": "NY",
//		},
//	}
//	options := DefaultOptions()
//	flattened := FlattenMap(data, options)
//	fmt.Println(flattened)
//
// Output:
//
//	map[address.city:New York address.state:NY age:30 name:John]
func FlattenMap(data map[string]interface{}, options Options) map[string]interface{} {
	flattened := make(map[string]interface{})
	flatten("", data, flattened, options.KeyDelimiter, options.MaxDepth)
	return flattened
}

// flatten is a helper function that recursively flattens a JSON object.
func flatten(prefix string, value interface{}, flattened map[string]interface{}, delimiter string, maxDepth int) {
	if maxDepth == 0 {
		flattened[strings.TrimSuffix(prefix, delimiter)] = value
		return
	}

	switch v := value.(type) {
	case map[string]interface{}:
		for key, val := range v {
			flatten(prefix+key+delimiter, val, flattened, delimiter, maxDepth-1)
		}
	case []interface{}:
		for i, val := range v {
			flatten(prefix+fmt.Sprintf("[%d]"+delimiter, i), val, flattened, delimiter, maxDepth-1)
		}
	default:
		flattened[strings.TrimSuffix(prefix, delimiter)] = v
	}
}

// UnflattenJSON unflattens a flattened JSON object into its original structure.
//
// Example:
//
//	flattened := map[string]interface{}{
//		"address.city":  "New York",
//		"address.state": "NY",
//		"age":           30,
//		"name":          "John",
//	}
//	options := DefaultOptions()
//	unflattened, err := UnflattenJSON(flattened, options)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	fmt.Println(unflattened)
//
// Output:
//
//	map[address:map[city:New York state:NY] age:30 name:John]
func UnflattenJSON(flattened map[string]interface{}, options Options) (interface{}, error) {
	result := make(map[string]interface{})
	for key, value := range flattened {
		setValue(result, strings.Split(key, options.KeyDelimiter), value)
	}
	return result, nil
}

// setValue is a helper function that sets a value in a nested map based on the given key path.
func setValue(data map[string]interface{}, keys []string, value interface{}) {
	lastKey := keys[len(keys)-1]
	parent := data
	for _, key := range keys[:len(keys)-1] {
		if _, ok := parent[key]; !ok {
			parent[key] = make(map[string]interface{})
		}
		parent = parent[key].(map[string]interface{})
	}
	parent[lastKey] = value
}
