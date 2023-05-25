package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 3, 5, 6}
	var target int = 5

	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return i
		}
	}

	return len(nums)
}
