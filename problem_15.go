package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	balance := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println(balance)
			break
		}

		data := strings.Split(scanner.Text(), " ")

		transactionType := data[0]
		dollars, _ := strconv.Atoi(data[1])

		if transactionType == "D" {
			balance += dollars
		}

		if transactionType == "W" {
			balance -= dollars
		}

	}

}
