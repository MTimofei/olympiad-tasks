package main

import (
	"fmt"
)

func main() {
	var digits = []int{9, 9}

	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	i := len(digits) - 1

	for {
		digits[i] += 1
		if digits[i] != 10 {
			return digits
		}
		if i == 0 {
			var digits2 = []int{1}
			digits2 = append(digits2, digits...)
			digits2[1] = 0
			return digits2
		}
		digits[i] = 0
		i -= 1
	}
}
