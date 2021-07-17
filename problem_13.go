package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
	Return total uppercase letters in a string
*/
func getTotalUppercase(inputString string) int {
	total := 0
	for _, letter := range inputString {
		if unicode.IsUpper(letter) {
			total += 1
		}
	}
	return total
}

/*
	Return total lowercase letters in a string
*/
func getTotalLowercase(inputString string) int {
	total := 0
	for _, letter := range inputString {
		if unicode.IsLower(letter) {
			total += 1
		}
	}
	return total
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		totalUppercase := getTotalUppercase(scanner.Text())
		totalLowercase := getTotalLowercase(scanner.Text())

		fmt.Println(totalUppercase)
		fmt.Println(totalLowercase)
	}

}
