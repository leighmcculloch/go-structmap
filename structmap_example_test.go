package structmap

import "fmt"

func ExampleMap() {
	s := struct {
		FieldA string
		FieldB int
	}{
		FieldA: "text",
		FieldB: 123,
	}

	m := Map(s)

	fmt.Println(m["FieldA"], m["FieldB"])

	// Output:
	// text 123
}

func ExampleStruct() {
	s := struct {
		FieldA string
		FieldB int
	}{}
	m := map[string]interface{}{
		"FieldA": "text",
		"FieldB": 123,
	}
	Struct(&s, m)

	fmt.Println(s)

	// Output:
	// {text 123}
}
