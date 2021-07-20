package main

import (
	"fmt"
	"strconv"
)

func addNumbers(x string, y string) int {
	xInt, _ := strconv.Atoi(x)
	yInt, _ := strconv.Atoi(y)
	return xInt + yInt
}

func main() {
	fmt.Println(addNumbers("3", "4"))
}
