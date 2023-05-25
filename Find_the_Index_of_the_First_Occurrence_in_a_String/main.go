package main

import (
	"fmt"
	"strings"
)

func main() {
	var haystack string = "leetcode"
	var needle string = "leeto"

	fmt.Println(strStr(haystack, needle))

}

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if ok := strings.HasPrefix(haystack, needle); ok {
			return i
		}
		haystack = haystack[1:]
	}
	return -1
}
