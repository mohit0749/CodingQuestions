package main

import (
	"fmt"
	"math"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
Time Complexity of this solution is O(k*n)
*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pairs := make([][]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return pairs
	}
	index := make([]int, len(nums1), len(nums1))
	k = Min(len(nums2)*len(nums1), k)
	for k > 0 {
		var minIndex = 0
		var minSum = math.MaxInt32
		for i, v := range nums1 {
			if index[i] < len(nums2) {
				currSum := v + nums2[index[i]]
				if currSum < minSum {
					minIndex = i
					minSum = currSum
				}
			}
		}
		pairs = append(pairs, []int{nums1[minIndex], nums2[index[minIndex]]})
		index[minIndex]++
		k--
	}
	return pairs
}

func main() {
	ar1 := []int{1, 1, 2}
	ar2 := []int{1, 2, 3}
	k := 10
	pairs := kSmallestPairs(ar1, ar2, k)
	for i, _ := range pairs {
		fmt.Println(pairs[i][0], pairs[i][1])
	}
}
