package main

func rob(nums []int) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp2 := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp2[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		var last2, last3, currMax int
		if i-2 >= 0 {
			last2 = dp[i-2]
		}
		if i-3 >= 0 {
			last3 = dp[i-3]
		}
		if last2 < last3 {
			currMax = last3
		} else {
			currMax = last2
		}
		dp[i] = nums[i] + currMax
	}

	for i := 2; i < len(nums); i++ {
		var last2, last3, currMax int
		if i-2 >= 0 {
			last2 = dp2[i-2]
		}
		if i-3 >= 0 {
			last3 = dp2[i-3]
		}
		if last2 < last3 {
			currMax = last3
		} else {
			currMax = last2
		}
		dp2[i] = nums[i] + currMax
	}

	var maxLoot int
	maxLoot = max(max(dp[len(dp)-2], dp[len(dp)-3]), max(dp2[len(dp2)-1], dp2[len(dp2)-2]))
	//if dp[len(nums)-2] > dp[len(nums)-1] {
	//	maxLoot = dp[len(nums)-2]
	//} else {
	//	if flag[len(nums)-1] == true {
	//		maxLoot = max(dp[len(nums)-1]-elementCnt[len(nums)-1]*nums[0], max(dp[len(nums)-1]-nums[len(nums)-1], dp[len(nums)-2]))
	//	} else {
	//		maxLoot = dp[len(nums)-1]
	//	}
	//}
	return maxLoot
}

func main() {
	//nums := []int{2, 3, 2, 2, 2, 2, 2, 2}
	nums := []int{200, 3, 140, 20, 10}
	//nums := []int{2, 2, 4, 3, 2, 5}
	println(rob(nums))
}
