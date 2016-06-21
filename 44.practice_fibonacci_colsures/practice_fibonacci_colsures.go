package main

import "fmt"

func fibonacci_iterator(x, a, b int) int {
  if x == 0 {
    return a
  }
  return fibonacci_iterator(x - 1, b, a + b)
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  n := 0

  return func() int {
    fibo := fibonacci_iterator(n, 0, 1)
    n += 1
    return fibo
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}

