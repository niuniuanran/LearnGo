package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1, f2, pos := 0, 1, -1
	return func() int {
		pos++
		if pos < 2 {
			return pos
		}
		nf2 := f1 + f2
		f1 = f2
		f2 = nf2
		return nf2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
