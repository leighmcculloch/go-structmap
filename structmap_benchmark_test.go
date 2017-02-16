package structmap

import (
	"testing"
)

type benchmarkStruct struct {
	A string
	B int
}

func BenchmarkMap(b *testing.B) {
	s := benchmarkStruct{
		A: "text",
		B: 123,
	}

	for i := 0; i < b.N; i++ {
		Map(s)
	}
}

func BenchmarkStruct(b *testing.B) {
	var s benchmarkStruct

	m := map[string]interface{}{
		"A": "text",
		"B": 123,
	}

	for i := 0; i < b.N; i++ {
		Struct(&s, m)
	}
}
