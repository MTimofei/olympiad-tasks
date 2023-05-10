package main

import (
	"fmt"
	"strings"
)

func main() {
	var i int = 10
	fmt.Println(isPalindrome(i))
}

func isPalindrome(x int) bool {
	var revers string = ""
	string := fmt.Sprintf("%d", x)

	chars := strings.Split(string, "")

	for i := len(chars) - 1; i >= 0; i-- {
		revers += chars[i]
	}
	if string == revers {
		return true
	}
	return false
}
