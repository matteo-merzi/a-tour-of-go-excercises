package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

var m map[string]int

func WordCount(s string) map[string]int {

	words := strings.Fields(s)

	m = make(map[string]int)

	for i := 0; i < len(words); i++ {
		word := words[i]
		elem := m[word]
		m[word] = elem + 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
