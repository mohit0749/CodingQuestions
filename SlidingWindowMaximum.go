package main

/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Follow up:
Could you solve it in linear time?

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/

type maxElement struct {
	Value, Count int
}

func maxSlidingWindow(nums []int, k int) []maxElement {
	ar := make([]maxElement, 0, len(nums)/k+1)
	if len(nums) == 0 && k == 0 {
		return ar
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	maxInCurrWindow := maxElement{
		Value: nums[0],
		Count: 1,
	}
	for i := 1; i < k; i++ {
		if nums[i] == maxInCurrWindow.Value {
			maxInCurrWindow.Count++
		}
		if nums[i] > maxInCurrWindow.Value {
			maxInCurrWindow.Value = nums[i]
			maxInCurrWindow.Count++
		}
	}
	ar[0] = maxInCurrWindow
	for i := k; i < len(nums)-k; i++ {
		if nums[i-k]==maxInCurrWindow.Value{
			maxInCurrWindow.Count--
		}
		temp:=max(maxInCurrWindow.Value,nums[i])
	}
}
