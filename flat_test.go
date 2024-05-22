package goflat_test

import (
	"reflect"
	"testing"

	goflat "github.com/brian-s-side-project/go-flat"
)

const (
	errorFlatteningJSON          = "Error flattening JSON: %+v"
	errorFlattenedJSONMismatch   = "Flattened JSON does not match expected result"
	addressStreetKey             = "address.street"
	addressCityKey               = "address.city"
	hobbies0Key                  = "hobbies.0"
	hobbies1Key                  = "hobbies.1"
	addressStreet                = "123 Main St"
	addressCity                  = "New York"
	hobbies0                     = "reading"
	hobbies1                     = "gaming"
	aBCD                         = "a.b.c.d"
	errorFlattenedMapMismatch    = "Flattened map does not match expected result"
	errorUnflatteningJSON        = "Error unflattening JSON: %+v"
	errorUnflattenedJSONMismatch = "Unflattened JSON does not match expected result"
)

func TestFlattenJSON(t *testing.T) {
	// Test case 1: Flattening a simple JSON object
	data := []byte(`{"name": "John", "age": 30}`)
	expected := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	options := goflat.DefaultOptions()
	result, err := goflat.FlattenJSON(data, options)
	if err != nil {
		t.Errorf(errorFlatteningJSON, err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorFlattenedJSONMismatch)
	}

	// Test case 2: Flattening a JSON object with nested objects and arrays
	data = []byte(`{"name": "John", "age": 30, "address": {"street": "123 Main St", "city": "New York"}, "hobbies": ["reading", "gaming"]}`)

	expected = map[string]interface{}{
		"name":           "John",
		"age":            30,
		addressStreetKey: addressStreet,
		addressCityKey:   addressCity,
		hobbies0Key:      hobbies0,
		hobbies1Key:      hobbies1,
	}
	result, err = goflat.FlattenJSON(data, options)
	if err != nil {
		t.Errorf(errorFlatteningJSON, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorFlattenedJSONMismatch)
	}

	// Test case 3: Flattening a JSON object with nested objects and arrays up to 4 levels deep
	data = []byte(`{"a": {"b": {"c": {"d": "value"}}}}`)
	expected = map[string]interface{}{
		aBCD: "value",
	}
	result, err = goflat.FlattenJSON(data, options)
	if err != nil {
		t.Errorf(errorFlatteningJSON, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorFlattenedJSONMismatch)
	}

	// test case 4: fail to unmarshal JSON
	data = []byte(`{"name": "John", "age": 30`)
	_, err = goflat.FlattenJSON(data, options)
	if err == nil {
		t.Errorf("Expected error when unmarshalling invalid JSON")
	}

	// test case 5: max depth is 0
	data = []byte(`{"a": {"b": {"c": {"d": "value"}}}}`)
	options.MaxDepth = 0
	expected = map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": map[string]interface{}{
					"d": "value",
				},
			},
		},
	}
	result, err = goflat.FlattenJSON(data, options)
	if err != nil {
		t.Errorf(errorFlatteningJSON, err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorFlattenedJSONMismatch)
	}
}

func TestFlattenMap(t *testing.T) {
	// Test case 1: Flattening a simple map
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	expected := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	options := goflat.DefaultOptions()
	result := goflat.FlattenMap(data, options)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorFlattenedMapMismatch)
	}

	// Test case 2: Flattening a map with nested maps and arrays
	data = map[string]interface{}{
		"name": "John",
		"age":  30,
		"address": map[string]interface{}{
			"street": addressStreet,
			"city":   addressCity,
		},
		"hobbies": []string{"reading", "gaming"},
	}
	expected = map[string]interface{}{
		"name":           "John",
		"age":            30,
		addressStreetKey: addressStreet,
		addressCityKey:   addressCity,
		hobbies0Key:      "reading",
		hobbies1Key:      "gaming",
	}
	result = goflat.FlattenMap(data, options)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorFlattenedMapMismatch)
	}

	// Test case 3: Flattening a map with nested maps up to 4 levels deep
	data = map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": map[string]interface{}{
					"d": "value",
				},
			},
		},
	}
	expected = map[string]interface{}{
		aBCD: "value",
	}
	result = goflat.FlattenMap(data, options)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorFlattenedMapMismatch)
	}
}

func TestUnflattenJSON(t *testing.T) {
	// Test case 1: Unflattening a simple flattened JSON object
	flattened := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	expected := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	options := goflat.DefaultOptions()
	result, err := goflat.UnflattenJSON(flattened, options)
	if err != nil {
		t.Errorf(errorUnflatteningJSON, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorUnflattenedJSONMismatch)
	}

	// Test case 2: Unflattening a flattened JSON object with nested keys
	flattened = map[string]interface{}{
		"name":           "John",
		"age":            30,
		addressStreetKey: addressStreet,
		addressCityKey:   addressCity,
		hobbies0Key:      "reading",
		hobbies1Key:      "gaming",
	}
	expected = map[string]interface{}{
		"name": "John",
		"age":  30,
		"address": map[string]interface{}{
			"street": addressStreet,
			"city":   addressCity,
		},
		"hobbies": []interface{}{"reading", "gaming"},
	}
	result, err = goflat.UnflattenJSON(flattened, options)
	if err != nil {
		t.Errorf(errorUnflatteningJSON, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorUnflattenedJSONMismatch)
	}

	// Test case 3: Unflattening a flattened JSON object with nested keys up to 4 levels deep
	flattened = map[string]interface{}{
		aBCD: "value",
	}
	expected = map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": map[string]interface{}{
					"d": "value",
				},
			},
		},
	}
	result, err = goflat.UnflattenJSON(flattened, options)
	if err != nil {
		t.Errorf(errorUnflatteningJSON, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errorUnflattenedJSONMismatch)
	}
}
