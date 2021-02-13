package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecur(t, ch)
	close(ch)
}

func walkRecur(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkRecur(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkRecur(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)

	for n1 := range ch1 {
		n2, ok := <-ch2
		if !ok {
			return false
		}
		if n1 != n2 {
			return false
		}
	}
	_, end := <-ch2
	return !end
}

func main() {
	fmt.Println(Same(tree.New(100), tree.New(2)))
}
