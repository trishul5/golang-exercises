package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortStrings(input string) string {
	words := strings.Split(input, ",")
	sort.Strings(words)
	return strings.Join(words, ",")
}

func main() {
	unsortedWords := "trishul,is,going,to,get,a,good,job"
	fmt.Println(sortStrings(unsortedWords))
}
