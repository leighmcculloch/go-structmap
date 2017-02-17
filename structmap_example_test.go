package structmap

import "fmt"

func ExampleMap() {
	s := struct {
		A string
		B int
	}{
		A: "text",
		B: 123,
	}

	m := Map(s)

	fmt.Println(m["A"], m["B"])

	// Output:
	// text 123
}

func ExampleStruct() {
	s := struct {
		A string
		B int
	}{}
	m := map[string]interface{}{
		"A": "text",
		"B": 123,
	}
	Struct(&s, m)

	fmt.Println(s)

	// Output:
	// {text 123}
}
