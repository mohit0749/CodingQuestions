package main

import (
	"fmt"
)

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}

	dp[0] = nums[0]
	if len(nums) >= 2 {
		dp[1] = nums[1]
		for i := 2; i < len(nums); i++ {
			a := dp[i-2] + nums[i]
			b := 0
			if i-3 >= 0 {
				b = dp[i-3] + nums[i]
			}
			if a > b {
				dp[i] = a
			} else {
				dp[i] = b
			}
		}

		if dp[len(dp)-1] >= dp[len(dp)-2] {
			return dp[len(dp)-1]
		}
		return dp[len(dp)-2]
	} else {
		return dp[0]
	}

}

func main() {
	nums := []int{2, 7, 2, 0, 2}
	fmt.Println(rob(nums))
}
