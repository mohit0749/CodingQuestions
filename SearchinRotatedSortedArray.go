package main

import "fmt"

func main() {
	ar := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Print(search(ar, 3))
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, low, high, target int) int {
	if high < low {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	} else if nums[low] == target {
		return low
	} else if nums[high] == target {
		return high
	}
	if nums[low] <= nums[mid] {
		if target < nums[mid] && target > nums[low] {
			return binarySearch(nums, low, mid-1, target)
		}
		return binarySearch(nums, mid+1, high, target)
	}
	if target > nums[mid] && target < nums[high] {
		return binarySearch(nums, mid+1, high, target)
	}
	return binarySearch(nums, low, mid-1, target)
}
