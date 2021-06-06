package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	fmt.Println(nums)
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	n := len(nums)
	i := n - 1
	for ; i >= 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	sort.Ints(nums[i:])
	fmt.Println(nums)

	fmt.Println(i - 1)
	for j, v := range nums[i:] {
		if v > nums[i-1] {
			fmt.Println(v, nums[i-1])
			fmt.Println(j)
			nums[i+j], nums[i-1] = nums[i-1], nums[i+j]
			break
		}
	}
	//sort.Ints(nums[i:])
}

func main() {
	nums := []int{2, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
