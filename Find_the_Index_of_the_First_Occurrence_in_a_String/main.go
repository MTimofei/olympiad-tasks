package main

import (
	"fmt"
	"strings"
)

func main() {
	var haystack string = "hello"
	var needle string = "ll"

	fmt.Println(strStr(haystack, needle))

}

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if ok := strings.HasPrefix(haystack, needle); ok {
			return i
		}
		if len(haystack) < len(needle) {
			return -1
		}
		haystack = haystack[1:]
	}
	return -1
}
