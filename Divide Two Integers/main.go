package main

import (
	"fmt"
)

func main() {
	var dividend int = -2147483648
	var divisor int = -1

	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	var result int = dividend / divisor
	if result >= 2147483648 {
		return 2147483647
	}

	return result

}
