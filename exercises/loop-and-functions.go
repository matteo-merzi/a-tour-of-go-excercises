package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	y := 1.0
	z := 0.0
	for {
		z = y
		y = y - (y*y-x)/2*y

		if math.Abs(y-z) < 1e-4 {
			break
		}
	}

	return y
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
