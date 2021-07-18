package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	results:=make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		results[i]=-1
	}
	n := len(nums)
	for i := 0; i < 2*n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			results[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return results
}

func main() {
	nums := []int{1, 3, 5, 2, 4}
	fmt.Println(nextGreaterElements(nums))
}
