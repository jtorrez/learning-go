package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  for n := 0; n < 10; n++ {
    z = z - ((z*z - x) / (2 * z))
  }
  return z
}

func main() {
  const x = 2
  mine, theirs := Sqrt(x), math.Sqrt(x)
  fmt.Println(mine, theirs, mine-theirs)
}
