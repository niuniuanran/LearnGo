package main

import (
	"fmt"
	"math/rand"
)

func split(a, b int) (x, y int) {
	x = a / b
	y = a - (x * b)
	return
}

func seedRandom() {
	rand.Seed(10)
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(1000))
}

func main() {
	var a, b = 14, 3
	fmt.Println(split(a, b))
}
