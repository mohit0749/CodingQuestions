package main

import (
	"container/list"
	"fmt"
)

/*
You are given an array of integers nums, there is a sliding window of size k which is
moving from the very left of the array to the very right. You can only see
the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.
*/

type keyValue struct {
	index int
	value int
}

func maxSlidingWindow(nums []int, k int) []int {
	maxElement := make([]int, 0)
	if len(nums) == 0 {
		return maxElement
	}

	dq := list.New()
	dq.PushBack(keyValue{index: 0, value: nums[0]})
	for i := 1; i < k; i++ {
		elem := dq.Back()
		for elem.Value.(keyValue).value < nums[i] {
			dq.Remove(elem)
			if dq.Len() == 0 {
				break
			}
			elem = dq.Back()
		}
		dq.PushBack(keyValue{index: i, value: nums[i]})
	}
	maxElement = append(maxElement, dq.Front().Value.(keyValue).value)
	for i := k; i < len(nums); i++ {
		elem := dq.Back()
		for elem.Value.(keyValue).value < nums[i] {
			dq.Remove(elem)
			if dq.Len() == 0 {
				break
			}
			elem = dq.Back()
		}
		dq.PushBack(keyValue{index: i, value: nums[i]})
		elem = dq.Front()
		for elem.Value.(keyValue).index <= i-k {
			dq.Remove(elem)
			if dq.Len() == 0 {
				break
			}
			elem = dq.Front()
		}
		maxElement = append(maxElement, dq.Front().Value.(keyValue).value)
	}
	return maxElement
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Printf("%+v\n", maxSlidingWindow(nums, k))
}
