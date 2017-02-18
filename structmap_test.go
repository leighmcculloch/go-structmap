package structmap

import (
	"reflect"
	"testing"
)

type testStruct struct {
	FieldA string
	FieldB int
}

func TestMap(t *testing.T) {
	tests := []struct {
		Struct      testStruct
		ExpectedMap map[string]interface{}
	}{
		{
			testStruct{FieldA: "text", FieldB: 123},
			map[string]interface{}{"FieldA": "text", "FieldB": 123},
		},
		{
			testStruct{FieldB: 123},
			map[string]interface{}{"FieldA": "", "FieldB": 123},
		},
		{
			testStruct{},
			map[string]interface{}{"FieldA": "", "FieldB": 0},
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
			map[string]interface{}{"FieldA": "text", "FieldB": 123},
			testStruct{FieldA: "text", FieldB: 123},
		},
		{
			map[string]interface{}{"FieldA": "", "FieldB": 123},
			testStruct{FieldB: 123},
		},
		{
			map[string]interface{}{"FieldA": "", "FieldB": 0},
			testStruct{},
		},
		{
			map[string]interface{}{"FieldB": 123},
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

func TestStructPanic(t *testing.T) {
	tests := []struct {
		Map           map[string]interface{}
		ExpectedPanic string
	}{
		{
			map[string]interface{}{"FieldC": ""},
			"Field FieldC does not exist on *structmap.testStruct",
		},
		{
			map[string]interface{}{"FieldA": "text", "FieldC": ""},
			"Field FieldC does not exist on *structmap.testStruct",
		},
	}

	for _, test := range tests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					if r.(error).Error() != test.ExpectedPanic {
						t.Errorf("Got panic %s, want %s", r, test.ExpectedPanic)
					}
				} else {
					t.Errorf("Did not panic, want %s", test.ExpectedPanic)
				}
			}()

			var s testStruct
			Struct(test.Map, &s)
		}()
	}
}
