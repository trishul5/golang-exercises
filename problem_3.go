package main

import "fmt"

func squaresMap(n int) map[int]int {

	squares := make(map[int]int)

	for i := 1; i <= n; i++ {
		squares[i] = i * i
	}

	return squares
}

func main() {
	squaresToNine := squaresMap(9)

	for i, square := range squaresToNine {
		fmt.Printf("%d ----> %d\n", i, square)
	}

}
