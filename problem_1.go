package main

import "fmt"

/*
	Write a program which will find al such numbers which are divisible by 7 but are not
	not a multiple of 5 between 2000 adn 3200 (both included)
*/

func validNumbers() {
	for i := 2000; i <= 3200; i++ {
		if (i%7 == 0) && (i%5 != 0) {
			fmt.Println(i)
		}
	}
}

func main() {
	validNumbers()
}
