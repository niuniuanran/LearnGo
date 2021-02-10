package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(1000))
}
