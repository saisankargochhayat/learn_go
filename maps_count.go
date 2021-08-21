package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	s_split := strings.Fields(s)
	m := make(map[string]int)
	for _, s_word := range s_split {
		m[s_word] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
