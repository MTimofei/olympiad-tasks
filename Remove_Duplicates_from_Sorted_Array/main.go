package main

import (
	"fmt"
)

func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	var n int = len(nums)
	var index []int

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			index = append(index, i)
		}
	}
	n = len(index) - 1
	for i := n; i > -1; i-- {

		copy(nums[index[i]:], nums[index[i]+1:])

		nums[len(nums)-1] = 0

		nums = nums[:len(nums)-1]
	}

	return len(nums)
}
