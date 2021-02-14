package main

import "sort"

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if k > len(nums) {
		k = len(nums)
	}
	left := 0
	for i := left; i < k && i < len(nums); i++ {
		for j := i + 1; j <= left+k && j < len(nums); j++ {
			var diff int
			if nums[i] < nums[j] {
				diff = nums[j] - nums[i]
			} else {
				diff = nums[i] - nums[j]
			}
			if diff <= t {
				return true
			}
		}
	}
	left++
	for i := k + 1; i < len(nums); i++ {
		for j := left; j < i; j++ {
			var diff int
			if nums[j] > nums[i] {
				diff = nums[j] - nums[i]
			} else {
				diff = nums[i] - nums[j]
			}
			if diff <= t {
				return true
			}
		}
		left++
	}
	return false
}

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	type elem struct {
		idx, val int
	}
	arr := make([]elem, 0)
	for i, v := range nums {
		arr = append(arr, elem{
			idx: i,
			val: v,
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].val <= arr[j].val && arr[i].idx <= arr[j].idx {
			return true
		}
		if arr[i].val < arr[j].val {
			return true
		}
		return false
	})
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			diff := arr[j].val - arr[i].val
			if diff > t {
				break
			}
			if abs(arr[i].idx-arr[j].idx) <= k {
				return true
			}
		}
	}
	return false
}


func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	//nums := []int{1, 2, 3, 1}
	//k := 3
	//t := 0

	//nums := []int{1, 0, 1, 1}
	//k := 1
	//t := 2

	//nums := []int{1, 5, 9, 1, 5, 9}
	//k := 3
	//t := 3

	nums := []int{2147483646, 2147483647}
	k := 3
	t := 3
	println(containsNearbyAlmostDuplicate(nums, k, t))
}
