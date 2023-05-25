package main

import (
	"fmt"
)

func main() {
	var nums = []int{3, 2, 2, 3}
	var val int = 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	var index []int

	for i := 0; i < len(nums); i++ {
		if val == nums[i] {
			index = append(index, i)
		}
	}

	n := len(index) - 1
	for i := n; i > -1; i-- {
		copy(nums[index[i]:], nums[index[i]+1:])
		nums[len(nums)-1] = 0
		nums = nums[:len(nums)-1]
	}

	return len(nums)
}
