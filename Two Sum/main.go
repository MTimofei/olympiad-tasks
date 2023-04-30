package main

func main() {

}

func twoSum(nums []int, target int) []int {
	count := len(nums)
	answer := []int{}
	for x := 0; x < count; x++ {
		for i := x + 1; i < count; i++ {
			if nums[x]+nums[i] == target {
				answer = append(answer, x)
				answer = append(answer, i)
				break
			}
		}
	}
	return answer
}
