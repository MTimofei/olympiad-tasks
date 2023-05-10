package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("MMDXLI"))
}

func romanToInt(s string) int {
	var wordbook = map[string]uint{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var numbers []int
	var ansver int

	chars := strings.Split(s, "")

	for _, i := range chars {
		numbers = append(numbers, int(wordbook[i]))
	}
	ansver = numbers[len(numbers)-1]
	for i := len(numbers) - 1; i > 0; i-- {

		if numbers[i] < numbers[i-1] {
			ansver += numbers[i-1]
			continue
		}
		if numbers[i] > numbers[i-1] {
			ansver -= numbers[i-1]
			continue
		}
		if numbers[i] == numbers[i-1] {
			ansver += numbers[i-1]
			continue
		}
	}

	return ansver
}
