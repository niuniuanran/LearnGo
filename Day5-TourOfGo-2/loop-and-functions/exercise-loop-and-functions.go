package main

import (
	"fmt"
)

// ErrNegativeSqrt is an error when trying to do Sqrt on negative number
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

// Sqrt calculates an estimation of the square root of the given float64 x
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	sqr := 1.0
	change := sqr
	for change*change > 1e-10 {
		change = (sqr*sqr - x) / (2 * sqr)
		sqr = sqr - change
	}
	return sqr, nil
}

func main() {
	var x float64
	fmt.Println("Please enter a number")
	fmt.Scanf("%f", &x)
	fmt.Println(Sqrt(x))
}
