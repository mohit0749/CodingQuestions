package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	l := 0
	for i := 0; i < len(nums); {
		nums[l] = nums[i]
		i++
		for i < len(nums) && nums[l] == nums[i] {
			i++
		}
		l += 1
	}
	fmt.Println(nums, l)
	return l
}
