package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	currSet := make([]int, 0)
	sort.Ints(candidates)
	getSumComb(candidates, 0, target, currSet, &ans)
	return ans
}

func getSumComb(candidates []int, start, target int, currSet []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, currSet)
		return
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		if start < i && candidates[i] == candidates[i-1] {
			continue
		}
		if target-candidates[i] < 0 {
			break
		}
		getSumComb(candidates, i+1, target-candidates[i], append(currSet, candidates[i]), ans)
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Printf("%+v\n", combinationSum2(candidates, target))
	candidates = []int{3, 1, 3, 5, 1, 1}
	target = 8

	fmt.Printf("%+v\n", combinationSum2(candidates, target))
}
