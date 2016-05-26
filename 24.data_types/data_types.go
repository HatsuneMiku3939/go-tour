package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe bool = false
  MaxInt uint64 = 1 << 64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

// Go data types
//
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64
//
// byte(same as uint8)
//
// rune(same as uint32, also represent Unicode code points)
//
// float32 float64
//
// complex64 complex128

func main() {
  const f = "%T(%v)\n"
  fmt.Printf(f, ToBe, ToBe)
  fmt.Printf(f, MaxInt, MaxInt)
  fmt.Printf(f, z, z)
}

