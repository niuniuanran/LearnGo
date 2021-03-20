package prime

import "math"

func IsPrime(num int) bool {
	if num < 2{
		return false
	}
	base := int(math.Sqrt(float64(num)))
	for i:= 2; i <=base; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}