package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	split_str := strings.Fields(s)
	word_map := make(map[string]int)
	for i := range split_str {
		word_map[split_str[i]] += 1
	}
	return word_map
}

func main() {
	wc.Test(WordCount)
}
