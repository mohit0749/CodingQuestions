package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	if k > n {
		k -= n
	}
	i := n - k - 1
	for ; i >= 0; i-- {
		for j := k; j >= 0; j-- {
			nums[i], nums[i+j] = nums[i+j], nums[i]
		}
	}
}

func main() {
	nums := []int{1, 2, 3}
	k := 5
	rotate(nums, k)
	fmt.Println(nums)
}
