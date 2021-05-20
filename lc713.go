package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] < k {
			return 1
		}
	}
	left, right := 0, 1
	currProd := nums[left]
	count := 0
	countNum := 0
	if nums[left] < k {
		countNum++
	}
	for left < right && right < len(nums) {
		if nums[right] < k {
			countNum++
		}
		currProd *= nums[right]

		if currProd < k {
			right++
			if right >= len(nums) {
				printWindow(left, right-1, nums)
				n := right - left
				count += (n * (n + 1)) / 2
				count -= n
				fmt.Println(currProd)

			}
		} else {
			printWindow(left, right-1, nums)
			n := right - left
			count += (n * (n + 1)) / 2
			count -= n

			for currProd >= k {
				currProd /= nums[left]
				left++
			}
			right++
			if right >= len(nums) {
				printWindow(left, right-1, nums)
				n := right - left
				count += (n * (n + 1)) / 2
				count -= n
				fmt.Println(currProd)

			}
		}
	}
	return count + countNum
}

func printWindow(l, r int, nums []int) {
	for i := l; i <= r; i++ {
		fmt.Print(nums[i], ",")
	}
	fmt.Println()
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 100
	println(numSubarrayProductLessThanK(nums, k))
}

[1,2,3,4,5]
[1,2,3,4]

