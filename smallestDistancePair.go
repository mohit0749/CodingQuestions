package main

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, m int) int {
	cnt := 0
	j := 0
	for i := 1; i < len(nums); i++ {
		for nums[i]-nums[j] > m && j < i {
			j += 1
		}
		cnt += i - j
	}

	return cnt
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < low {
			low = diff
		}
	}
	high := nums[len(nums)-1] - nums[0]
	fmt.Println(low, ",", high)
	for low < high {
		mid := (low + high) >> 1
		if findPairs(nums, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	ar := []int{38, 33, 57, 65, 13, 2, 86, 75, 4, 56}
	k := 26
	fmt.Println(smallestDistancePair(ar, k))
}
