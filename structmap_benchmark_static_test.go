package structmap

import (
	"testing"
)

type benchmarkStaticStruct struct {
	A string
	B int
}

func (s *benchmarkStaticStruct) Map() map[string]interface{} {
	return map[string]interface{}{
		"A": s.A,
		"B": s.B,
	}
}

func newWithMap(m map[string]interface{}) *benchmarkStaticStruct {
	return &benchmarkStaticStruct{
		A: m["A"].(string),
		B: m["B"].(int),
	}
}

func BenchmarkStaticMap(b *testing.B) {
	s := benchmarkStaticStruct{
		A: "text",
		B: 123,
	}

	for i := 0; i < b.N; i++ {
		s.Map()
	}
}

func BenchmarkStaticStruct(b *testing.B) {
	m := map[string]interface{}{
		"A": "text",
		"B": 123,
	}

	for i := 0; i < b.N; i++ {
		newWithMap(m)
	}
}
