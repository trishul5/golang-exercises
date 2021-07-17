// package main

// import (
// 	"fmt"
// )

// /*
// 	Function that returns true if all digits are even or false if not
// */
// func hasOnlyEvenDigits(n int) bool {

// 	flag := true

// 	for {
// 		if n == 0 {
// 			break
// 		} else {
// 			digit := n % 10

// 			if digit%2 != 0 {
// 				flag = false
// 			}

// 			n = n / 10
// 		}
// 	}
// 	return flag
// }

// func printNumbersWithEvenDigits(start int, end int) {
// 	for i := start; i <= end; i++ {
// 		if hasOnlyEvenDigits(i) {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main() {
// 	printNumbersWithEvenDigits(1000, 3000)
// }
