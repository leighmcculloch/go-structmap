// Package structmap converts structs to maps, and maps to structs.
package structmap

import (
	"reflect"

	"4d63.com/iszero"
	"4d63.com/structtags"
)

const tagKey = "structmap"

// Map converts a struct to a map of field names to values.
func Map(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)

		var key string
		var options structtags.TagOptions
		if tagStr, ok := f.Tag.Lookup(tagKey); ok {
			tag := structtags.Parse(tagStr)
			key = tag.Value
			options = tag.Options
		}

		if key == "" {
			key = f.Name
		}

		if options.Contains("omitempty") && iszero.Value(fv) {
			continue
		}

		value := fv.Interface()
		if f.Type.Kind() == reflect.Struct {
			value = Map(value)
		}

		m[key] = value
	}

	return m
}

// Struct fills a struct with the values in the map.
func Struct(m map[string]interface{}, s interface{}) {
	v := reflect.ValueOf(s)

	e := v.Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		f := t.Field(i)
		fv := e.Field(i)

		var key string
		if tagValue, ok := f.Tag.Lookup(tagKey); ok {
			key = structtags.Parse(tagValue).Value
		} else {
			key = f.Name
		}

		if value, ok := m[key]; ok {
			if f.Type.Kind() == reflect.Struct {
				Struct(value.(map[string]interface{}), fv.Addr().Interface())
			} else {
				fv.Set(reflect.ValueOf(value))
			}
		}
	}
}
