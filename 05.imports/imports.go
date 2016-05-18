package main

import (
  "fmt"
	"math"
)

func main() {
	fmt.Println("Now you have %g probles.",
		math.Nextafter(2, 3))
}
