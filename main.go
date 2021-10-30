package main

import (
	"math"
)

func f1(f float64) float64 {
	return 2 * f + 10
}
func f2(f float64) float64 {
	return 3.0 * float64(math.Pow(float64(f), 3))
}

func main() {

	Solve(f1, f2, -1000, 1000)
}
