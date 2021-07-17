package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
	Returns the count of digits in a string
*/
func getDigitCount(inputString string) int {
	total := 0
	for i := 0; i < len(inputString); i++ {
		if unicode.IsDigit(rune(inputString[i])) {
			total += 1
		}
	}
	return total
}

/*
	Returns the count of lettes in a string
*/

func getLetterCount(inputString string) int {
	total := 0
	for i := 0; i < len(inputString); i++ {
		if unicode.IsLetter(rune(inputString[i])) {
			total += 1
		}
	}
	return total
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	inputString := scanner.Text()
	totalDigits := getDigitCount(inputString)
	totalLetters := getLetterCount(inputString)

	fmt.Println(totalDigits)
	fmt.Println(totalLetters)

	fmt.Printf("LETTERS %d\n", totalLetters)
	fmt.Printf("DIGITS %d\n", totalDigits)
}
