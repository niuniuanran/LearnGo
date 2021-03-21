package lib

import "math"

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	base := int(math.Sqrt(float64(num)))
	for i := 2; i <= base; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Factors(num int) []int {
	if num < 0 {
		return append(Factors(num*-1), -1)
	}
	if num < 2 {
		return []int{num}
	}
	var results []int
	curr := num
	for curr > 1 {
		for i := 2; i <= num; i++ {
			for curr%i == 0 {
				curr = curr / i
				results = append(results, i)
			}
		}
	}
	return results
}
