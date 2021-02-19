package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[2:]
	lines := make(map[string]map[string]int)
	if len(files) < 1 {
		return
	}
	for _, f := range files {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if lines[line] == nil {
				lines[line] = make(map[string]int)
			}
			lines[line][f]++
		}
	}
	for s, line := range lines {
		if totalCount(line) > 1 {
			fmt.Printf("%3d\t%q\n", totalCount(line), s)
			for fname, n := range line {
				fmt.Printf("\t%2d\t%s\n", n, fname)
			}
		}
	}
}

func totalCount(lines map[string]int) int {
	sum := 0
	for _, count := range lines {
		sum += count
	}
	return sum
}
