// implement sqrt by newton's method with loop and function

package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  const damp = 0.000000000000001
  z := x / 2

  for i := 0; i < 10; i++ {
    z_new := z - (z * z - x) / (2 * z)
    diff := math.Abs(z_new - z)
    z = z_new

    if diff < damp {
      fmt.Println(i)
      return z
    }
  }

  return z
}

func main() {
  for i := 1.0; i < 101.0; i++ {
    std_sqrt := math.Sqrt(i)
    my_sqrt := Sqrt(i)
    fmt.Println(std_sqrt, my_sqrt, std_sqrt - my_sqrt)
  }
}
