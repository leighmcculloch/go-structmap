package structmap

import (
	"reflect"
	"testing"
)

type testStruct struct {
	FieldA string
	FieldB int `structmap:"b"`
	FieldC string `structmap:"c,omitempty"`
	FieldD string `structmap:",omitempty"`
}

func TestMap(t *testing.T) {
	tests := []struct {
		Struct      testStruct
		ExpectedMap map[string]interface{}
	}{
		{
			testStruct{FieldA: "text", FieldB: 123, FieldC: "123", FieldD: "456"},
			map[string]interface{}{"FieldA": "text", "b": 123, "c": "123", "FieldD": "456"},
		},
		{
			testStruct{FieldA: "text", FieldB: 123, FieldC: "123"},
			map[string]interface{}{"FieldA": "text", "b": 123, "c": "123"},
		},
		{
			testStruct{FieldB: 123},
			map[string]interface{}{"FieldA": "", "b": 123},
		},
		{
			testStruct{},
			map[string]interface{}{"FieldA": "", "b": 0},
		},
	}

	for _, test := range tests {
		m := Map(test.Struct)
		if !reflect.DeepEqual(test.ExpectedMap, m) {
			t.Errorf("Struct %v got %v, want %v", test.Struct, m, test.ExpectedMap)
		}
	}
}

func TestStruct(t *testing.T) {
	tests := []struct {
		Map            map[string]interface{}
		ExpectedStruct testStruct
	}{
		{
			map[string]interface{}{"FieldA": "text", "b": 123},
			testStruct{FieldA: "text", FieldB: 123},
		},
		{
			map[string]interface{}{"FieldA": "", "b": 123},
			testStruct{FieldB: 123},
		},
		{
			map[string]interface{}{"FieldA": "", "b": 0},
			testStruct{},
		},
		{
			map[string]interface{}{"b": 123},
			testStruct{FieldB: 123},
		},
		{
			map[string]interface{}{},
			testStruct{},
		},
	}

	for _, test := range tests {

		var s testStruct
		Struct(test.Map, &s)
		if test.ExpectedStruct != s {
			t.Errorf("Struct %v got %v, want %v", test.Map, s, test.ExpectedStruct)
		}
	}
}

func TestStructIgnoreMapKeysNotInStruct(t *testing.T) {
	tests := []struct {
		Map           map[string]interface{}
		ExpectedStruct testStruct
	}{
		{
			map[string]interface{}{"FieldE": ""},
			testStruct{},
		},
		{
			map[string]interface{}{"FieldA": "text", "FieldE": ""},
			testStruct{FieldA:"text"},
		},
		{
			map[string]interface{}{"FieldB": 123},
			testStruct{},
		},
	}

	for _, test := range tests {
		var s testStruct
		Struct(test.Map, &s)
		if test.ExpectedStruct != s {
			t.Errorf("Struct %v got %v, want %v", test.Map, s, test.ExpectedStruct)
		}
	}
}
