package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getOddNumbers(inputNumbers string) string {
	var oddNumbers []string
	numbers := strings.Split(inputNumbers, ",")

	for _, number := range numbers {
		convertedNumber, _ := strconv.Atoi(number)

		if convertedNumber%2 != 0 {
			convertedNumber := strconv.Itoa(convertedNumber)
			oddNumbers = append(oddNumbers, convertedNumber)
		}
	}
	return strings.Join(oddNumbers, ",")
}

func main() {

	fmt.Println(getOddNumbers("1,2,3,4,5,6,7,8,9,10"))

}
