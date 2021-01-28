package main

import (
	"fmt"
	"sort"
)

func minPlatforms(arr, dep []int) int {
	sort.Ints(arr)
	sort.Ints(dep)
	count := 0
	minCount := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, j := 0, 0; i < len(arr) && j < len(dep); {
		if arr[i] < dep[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		minCount = max(minCount, count)
	}
	return minCount
}

func main() {
	arr := []int{900, 1100, 1235}
	dep := []int{1000, 1200, 1240}
	fmt.Printf("%+v\n", minPlatforms(arr, dep))
}
