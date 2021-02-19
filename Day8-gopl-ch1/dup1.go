package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[2:]
	counts := make(map[string]int)
	if len(files) < 1 {
		return
	}
	for _, f := range files {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}

}
