package main

import "fmt"

func subsets(nums []int) [][]int {
	results := make([][]int, 0)
	generatePowerSet(nums, make([]int, 0), 0, &results)
	return results
}

func generatePowerSet(nums []int, pset []int, start int, results *[][]int) {
	*results = append(*results, pset)
	for i := start; i < len(nums); i++ {
		generatePowerSet(nums, append([]int{}, append(pset, nums[i])...), i+1, results)
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("%+v", subsets(nums))
}
