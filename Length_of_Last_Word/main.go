package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "luffy is still joyboy"

	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	return len(ss[len(ss)-1])
}
