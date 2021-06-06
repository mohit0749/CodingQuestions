package main

import "fmt"

func permute(nums []int) [][]int {
	permutations := make([][]int, 0)
	generatePermutation(len(nums), nums, make([]int, 0), &permutations)
	return permutations
}

func generatePermutation(n int, nums []int, permute []int, permutations *[][]int) {
	if len(permute) == n {
		*permutations = append(*permutations, permute)
		return
	}
	for i := 0; i < len(nums); i++ {
		newPermute := make([]int, len(permute))
		copy(newPermute, permute)
		newPermute = append(newPermute, nums[i])
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums = append(newNums[:i], newNums[i+1:]...)
		generatePermutation(n, newNums, newPermute, permutations)
	}
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Printf("%+v", permute(arr))
}
