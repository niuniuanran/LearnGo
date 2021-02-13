package main

import (
	"fmt"
	"math"
)

// Sqrt calculates an estimation of the square root of the given float64 x
func Sqrt(x float64) float64 {
	sqr := 1.0
	change := sqr
	for change*change > 1e-10 {
		change = (sqr*sqr - x) / (2 * sqr)
		sqr = sqr - change
	}
	return sqr
}

func main() {
	var x float64
	fmt.Println("Please enter a number")
	fmt.Scanf("%f", &x)
	fmt.Printf("My sqrt: %16f\n", Sqrt(x))
	fmt.Printf("math lib: %15f\n", math.Sqrt(x))
}
