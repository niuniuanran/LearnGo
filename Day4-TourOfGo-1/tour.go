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

func typeConvertion() (c complex128) {
	c = 32
	return
}

func myFirstLoop() {
	for sum := 1; sum < 1000; {
		sum += sum
		fmt.Println(sum)
	}
}

func main() {
	var a, b = 14, 3
	fmt.Println(split(a, b))
	fmt.Println(typeConvertion())
	myFirstLoop()
}
