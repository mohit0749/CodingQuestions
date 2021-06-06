package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	var result int
	backtrack(nums, 0, target, 0, &result, "")
	return result
}

func backtrack(nums []int, start, target int, sum int, result *int, expr string) {
	if start == len(nums) {
		fmt.Println(nums, sum, expr)
		if sum == target {
			*result += 1
		}
		return
	}
	backtrack(nums, start+1, target, sum+nums[start], result, expr+fmt.Sprintf("+%d", nums[start]))
	backtrack(nums, start+1, target, sum-nums[start], result, expr+fmt.Sprintf("-%d", nums[start]))
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	k := 3
	println(findTargetSumWays(nums, k))
}
