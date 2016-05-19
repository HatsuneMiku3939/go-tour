package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  // belows are same
  // var a string
  // var b string
  // a, b = swap("hello", "world")

  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
