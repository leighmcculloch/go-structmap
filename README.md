# structmap

[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/go-structmap)](https://goreportcard.com/report/github.com/leighmcculloch/go-structmap)
[![Codecov](https://img.shields.io/codecov/c/github/leighmcculloch/go-structmap.svg)](https://codecov.io/gh/leighmcculloch/go-structmap)
[![Build Status](https://img.shields.io/travis/leighmcculloch/go-structmap.svg)](https://travis-ci.org/leighmcculloch/go-structmap)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/leighmcculloch/go-structmap)

A go package for converting structs to maps, and maps to structs.

I mostly wrote this for fun and to learn how to use reflection in Go. I don't recommend using this package, because like it and similar libraries it is a complex solution where a simpler one can be used instead. See the [why not](#why-not) section below.

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
## Why

Inspired by [fatih/structs](https://github.com/fatih/structs) and [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure), but focuses purely on creating an exact one-to-one mapping of a struct and a map without any additional weight of tags, hooks, etc.

## Why not

There is a simpler way to solve this problem in many cases:

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
