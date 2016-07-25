package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  for x2 := 0.0 ; math.Abs(x2 - x) < 0.0001; {
    x2 = x - (x * x -2) / (x * 2)

    x = x2
  }

  return x
}

func main() {
  fmt.Println(Sqrt(5.0))
}
