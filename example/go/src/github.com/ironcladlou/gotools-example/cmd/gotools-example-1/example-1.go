package main

import (
  "fmt"

  "github.com/ironcladlou/gotools-example/lib"
)

func main() {
  foo := lib.Foo{
    StringField: "a string",
  }
  fmt.Println("command 1: %v", foo)
}
