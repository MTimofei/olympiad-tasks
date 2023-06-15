package main

import (
	"fmt"
)

func main() {

	n := 10

	fmt.Println(mySqrt(n))
}

func mySqrt(x int) int {
	switch {
	case x == 0:
		return 0
	case x < 4:
		return 1
	}

	var i int = 3

	for {
		switch {
		case x == i*i:
			return i
		case x < i*i:
			i--
			return i
		}
		i++
	}

}
