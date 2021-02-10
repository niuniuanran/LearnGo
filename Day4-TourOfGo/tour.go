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
	fmt.Println(split(12, 3))
}
