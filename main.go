package main

import (
	"fmt"
	"math"
)

func f1(f float64) float64 {
	return 2*f + 10
}
func f2(f float64) float64 {
	return 3.0 * float64(math.Pow(f, 3))
}

func main() {
	fmt.Println(Solve(f1, f2, -100, 100))
}
