package main

import "fmt"

import (
	"reflect"
)

var c int
var b, d bool

func test() {}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func main() {
  // var b int
  var i int
  var s string
  fmt.Println(reflect.TypeOf(i).String())
  fmt.Println(i)
  fmt.Println(len(s))
  fmt.Println(c)
  fmt.Println(swap("first", "second"))
  fmt.Println(add(100,200))
}
