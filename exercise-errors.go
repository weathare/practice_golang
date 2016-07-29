package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	for x2 := 1.0; math.Abs(x2-x) < 0.00000000000001; {
		x2 = x - (x*x-2)/(x*2)

		x = x2
	}

	return x, nil
}

func main() {
	fmt.Println(Sqrt(5.0))
	fmt.Println(Sqrt(-2))
}
