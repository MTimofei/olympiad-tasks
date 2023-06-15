package main

import (
	"fmt"
)

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	var n1 int = 1
	var n2 int = 2
	var out int
	switch {
	case n == 1:
		return 1
	case n == 2:
		return 2
	}

	for i := 2; i < n; i++ {
		out = n1 + n2
		n1, n2 = n2, out
	}
	return out
}
