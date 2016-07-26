package main

import (
  "fmt"
  "math"
)


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
  const delta = 1e-10

  if (x < 0) {
    return 0, ErrNegativeSqrt(x)
  }

  z := x / 2
  for i := 0; i < 10; i++ {
    z_new := z - (z * z - x) / (2 * z)
    diff := math.Abs(z_new - z)
    z = z_new

    if diff < delta {
      break
    }
  }

  return z, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
