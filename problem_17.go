package main

import (
	"fmt"
	"strings"
)

func countWords(inputString string) map[string]int {
	wordCounts := make(map[string]int)

	words := strings.Split(inputString, " ")

	for _, word := range words {
		wordCounts[word] += 1
	}

	return wordCounts
}

func main() {
	counts := countWords("foo bar baz bar")

	for word, count := range counts {
		fmt.Printf("%s -----> %d\n", word, count)
	}

}
