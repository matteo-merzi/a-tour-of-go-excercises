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
	if x <= 0 {
		return 0, ErrNegativeSqrt(x)
	}
	y := 1.0
	z := 0.0
	for {
		z = y
		y = y - (y*y-x)/2*y

		if math.Abs(y-z) < 0.0001 {
			break
		}
	}

	return y, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
