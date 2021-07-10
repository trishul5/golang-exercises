package main

import (
	"fmt"
	"math"
)

/*
	Question:
	Write a program that calculates and prints the value according to the given formula:
	Q = Square root of [(2 * C * D)/H]
	Following are the fixed values of C and H:
	C is 50. H is 30.
	D is the variable whose values should be input to your program in a comma-separated sequence.
	Example
	Let us assume the following comma separated input sequence is given to the program:
	100,150,180
	The output of the program should be:
	18,22,24
*/

const C = 50
const H = 30

func calculateValue(D float64) float64 {
	output := math.Sqrt((2 * C * D) / H)
	return output
}

func main() {
	answer := math.Round(calculateValue(100))
	fmt.Println(answer)
}
