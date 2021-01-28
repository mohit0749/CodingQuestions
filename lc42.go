package main

//trapping rain water problem
/*
Time Complexity : O(n)
Space Complexity: O(1)
*/
func trappedWater(height []int) int {
	l := 0
	h := len(height) - 1
	for height[l] == 0 {
		l++
	}
	for height[h] == 0 {
		h--
	}
	leftMax := 0
	rightMax := 0
	ans := 0
	for l < h {
		if height[l] < height[h] {
			if leftMax < height[l] {
				leftMax = height[l]
			} else {
				ans += leftMax - height[l]
			}
			l++
		} else {
			if rightMax < height[h] {
				rightMax = height[h]
			} else {
				ans += rightMax - height[h]
			}
			h--
		}
	}
	return ans
}

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(trappedWater(arr))
}
