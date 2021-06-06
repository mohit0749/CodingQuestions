package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int)
	results := make([]int, 0)
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{0, nums2[0]})
	for len(stack) > 0 {
		point := stack[len(stack)-1]
		j := point[0] + 1
		greater := -1
		for j < len(nums2) {
			point := stack[len(stack)-1]
			if point[1] < nums2[j] {
				greater = nums2[j]
				break
			}
			stack = append(stack, [2]int{j, nums2[j]})
			j++
		}
		if greater != -1 {
			for len(stack) > 0 {
				point := stack[len(stack)-1]
				if point[1] > greater {
					break
				}
				nextGreater[point[1]] = greater
				stack = stack[:len(stack)-1]
			}
		}
		if j == len(nums2) {
			break
		} else {
			stack = append(stack, [2]int{j, nums2[j]})
		}
	}
	for _, v := range nums1 {
		ele, ok := nextGreater[v]
		if !ok {
			ele = -1
		}
		results = append(results, ele)
	}
	return results
}

func main() {
	//[1,3,5,2,4]
	//	[6,5,4,3,2,1,7]
	nums1 := []int{1, 3, 5, 2, 4}
	nums2 := []int{6, 5, 4, 3, 2, 1, 7}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
