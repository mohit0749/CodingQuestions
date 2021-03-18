package main

import "fmt"

func findDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	n := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		nums[(nums[i]-1)%n] += n
	}
	fmt.Printf("%+v\n", nums)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		var cnt int = nums[i] / n
		if cnt >= 2 {
			ans = append(ans, 1+i)
		}
	}
	return ans
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// nums = []int{1, 2, 3, 4, 5, 6, 6, 7, 7}
	fmt.Printf("%+v\n", findDuplicates(nums))
}
