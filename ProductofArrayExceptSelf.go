package main

import "fmt"

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	prod := 1
	for i, v := range nums {
		prod *= v
		output[i] = prod
	}
	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i-1 >= 0 {
			output[i] = output[i-1] * prod
		} else {
			output[i] = prod
		}
		prod *= nums[i]
	}
	return output

}

func main() {
	ar := []int{1, 2, 3, 4}
	fmt.Print(productExceptSelf(ar))
}
