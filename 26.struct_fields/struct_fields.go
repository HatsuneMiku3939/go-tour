package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  // var v Vertex
  // v = Vertex{1, 2}
  v := Vertex{1, 2}

  v.X = 4
  fmt.Println(v)
}
