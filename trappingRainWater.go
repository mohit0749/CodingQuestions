package main

import "fmt"

func trap(height []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	lMax := 0
	rMax := len(height)-1
	low, high := 0, len(height)-1
	result := 0
	for low<len(height) && height[low]==0{
		low++
	}
	for high>=0 && height[high]==0{
		high--
	}
	lMax=low
	rMax=high
	for low < high {
		if height[lMax] < height[low] {
			lMax = low
		}
		if height[rMax] < height[high] {
			rMax = high
		}
		if height[low] <= height[high] {

			result += min(height[rMax], height[lMax]) - height[low]
			low++
		} else {

			result += min(height[rMax], height[lMax]) - height[high]
			high--
		}
	}
	return result
}

func main() {
   ar:=[]int{4,2,0,3,2,5}
   fmt.Println(trap(ar))
}
