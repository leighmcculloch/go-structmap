// Package structmap converts structs to maps, and maps to structs.
package structmap

import (
	"fmt"
	"reflect"
)

// Map converts a struct to a map of field names to values.
func Map(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return m
}

// Struct fills a struct with the values in the map.
func Struct(s interface{}, m map[string]interface{}) {
	v := reflect.ValueOf(s)

	for mk, mv := range m {
		f := v.Elem().FieldByName(mk)
		if !f.IsValid() {
			panic(fmt.Errorf("Field %s does not exist on %s", mk, v.Type()))
		}
		f.Set(reflect.ValueOf(mv))
	}
}
