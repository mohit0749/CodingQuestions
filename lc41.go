package main

//O(n) complexity + O(n) space
func _firstMissingPositive(nums []int) int {
	mp := make(map[int]bool)
	max := nums[0]
	for _, i := range nums {
		if i <= 0 {
			continue
		}
		mp[i] = true
		if max < i {
			max = i
		}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := mp[i]; ok {
			continue
		}
		return i
	}
	return max + 1
}

//O(n) complexity + O(1) space
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	nums = append(nums, -1)
	for i := 0; i < len(nums); i++ {
		_shiftToCorrectPosition(nums, i)
	}

	for i := 1; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums) + 1
}

func _shiftToCorrectPosition(nums []int, n int) {
	for nums[n] < len(nums) && nums[n] >= 0 && nums[n] != n {
		tmp := nums[nums[n]]
		nums[nums[n]] = nums[n]
		if tmp == nums[n] {
			nums[n] = -1
		} else {
			nums[n] = tmp
		}
	}
}

func main() {
	nums := []int{1, 2, 0}
	println(firstMissingPositive(nums))
	//fmt.Printf("%+v\n", nums)
}
