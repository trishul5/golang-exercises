package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	Returns capitalized words in a string with all lowercase words
*/
func capitalizeWords(inputText string) string {
	var capitalizedWords []string

	words := strings.Split(inputText, " ")

	fmt.Println(words)

	for _, word := range words {
		capitalizedWords = append(capitalizedWords, strings.ToUpper(word))
	}

	return strings.Join(capitalizedWords, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		fmt.Println(capitalizeWords(scanner.Text()))
	}
}
