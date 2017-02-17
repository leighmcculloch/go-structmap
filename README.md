# structmap
A go package for converting structs to maps, and maps to structs.

## Usage

```
go get github.com/leighmcculloch/go-structmap
```

```go
import "github.com/leighmcculloch/go-structmap"

type S {
  A string
  B int
}

func main() {
  s := S{
    A: "text",
    B: 123,
  }

  m := structmap.Map(s)
  // m is map[A:text B:123]

  var s2 S
  structmap.Struct(&s2, m)
  // s2 is {text 123}
}
```

## Why not to use

It is more performant and the compiler can provide more safety checks if you do the following instead.

```go
type S {
  A string
  B int
}

func (s *S) Map() map[string]interface{} {
	return map[string]interface{}{
		"A": s.A,
		"B": s.B,
	}
}

func NewWithMap(m map[string]interface{}) *S {
	return &S{
		A: m["A"].(string),
		B: m["B"].(int),
	}
}

func main() {
  s := S{
    A: "text",
    B: 123,
  }

  m := s.Map()
  // m is map[A:text B:123]

  s2 := NewWithMap(m)
  // s2 is {text 123}
}
```
