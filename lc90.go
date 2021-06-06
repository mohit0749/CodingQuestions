package main

import (
	"fmt"
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {
	results := make([][]int, 0)
	backtrack(nums, make([]int, 0), &results, make(map[string]bool))
	return results
}

func backtrack(nums []int, pset []int, results *[][]int, hash map[string]bool) {
	sort.Ints(pset)
	s := ""
	for _, v := range pset {
		s += "_" + strconv.Itoa(v)
	}
	fmt.Println(s)
	if !hash[s] {
		*results = append(*results, pset)
		hash[s] = true
	}
	mp := make(map[int]bool)
	for i, v := range nums {
		if !mp[v] {
			backtrack(append([]int{}, nums[i+1:]...), append([]int{}, append(pset, v)...), results, hash)
		}
		mp[v] = true
	}
}

func main() {
	nums := []int{4, 4, 4, 1, 4}
	subsetsWithDup(nums)
}
