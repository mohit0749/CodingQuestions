package main

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		var mid int = (l + r) / 2
		if mid%2 == 1 {
			mid++
		}
		if mid-1 >= 0 && nums[mid-1] == nums[mid] {
			r = mid - 2
		} else if mid+1 < len(nums) && nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			return nums[mid]
		}

	}
	return -1
}

func main() {
	//nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	//nums := []int{3,3,7,7,10,11,11}
	//nums:=[]int{2,2,1}
	nums:=[]int{}
	println(singleNonDuplicate(nums))
}
