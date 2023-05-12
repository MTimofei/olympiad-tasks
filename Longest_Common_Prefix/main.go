package main

import (
	"fmt"
	"strings"
)

func main() {
	var strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	var ferst []string
	var answer string
	var c int
	var ok bool = true
	if strs[0] == "" {
		return strs[0]
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs)-1; i++ {
		for j := 0; j < len(strs)-i-1; j++ {
			if len(strs[j]) > len(strs[j+1]) {
				strs[j+1], strs[j] = strs[j], strs[j+1]
			}
		}
	}

	ferst = strings.Split(strs[0], "")
	for c = 0; c < len(ferst) && ok; c++ {
		for j := 1; j < len(strs); j++ {
			s := strings.Split(strs[j], "")
			if ferst[c] != s[c] {
				ok = false
				break
			}
		}
		if ok {
			answer += ferst[c]
		}
	}
	return answer
}
