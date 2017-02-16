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
