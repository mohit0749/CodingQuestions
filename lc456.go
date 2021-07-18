package main

import "math"

func find132pattern(nums []int) bool {
	stack := make([]int, 0)
	s3 := math.MinInt32
	for _, v := range nums {
		if v < s3 {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < v {
			s3 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	return false
}

func main() {
	nums := []int{-2, 1, 2, -2, 1, 2} //=>{-2,2,1}
	println(find132pattern(nums))
}
