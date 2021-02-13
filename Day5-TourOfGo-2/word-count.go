package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount counts the frequency of each character in s
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		_, ok := m[w]
		if ok {
			m[w]++
		} else {
			m[w] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
