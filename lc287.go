package main

//O(n) time + O(1) time complexity
// using array as hash table to store the frequency
func findDuplicate1(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		nums[(nums[i]%n)-1] += n
	}
	for i := 0; i < len(nums); i++ {
		var tmp int = nums[i] / n
		if tmp > 1 {
			return i + 1
		}
	}
	return 0
}

//O(n) time + O(1) time complexity
// without modifying the array
// hare and tortise approach
func findDuplicate(nums []int) int {
	hare := nums[0]
	tortise := nums[0]
	tortise = nums[tortise]
	hare = nums[nums[hare]]
	for tortise != hare {
		tortise = nums[tortise]
		hare = nums[nums[hare]]
	}
	tortise = nums[0]
	for tortise != hare {
		tortise = nums[tortise]
		hare = nums[hare]
	}
	return hare
}

//it works only if duplicate appears two times only
//for case [2,2,2,2,2] it will fail
func findDuplicateXOR(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	for i := 1; i < len(nums); i++ {
		ans = ans ^ i
	}
	return ans
}

func main() {
	nums := []int{1, 3, 4, 4, 2}

	print(findDuplicate(nums))
}
