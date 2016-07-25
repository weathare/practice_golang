package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	x := make(map[string]int)
	for _, j := range strings.Fields(s) {
		x[j]++
	}

	return x
}

func main() {
	wc.Test(WordCount)
}
