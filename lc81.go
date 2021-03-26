package main

func search(nums []int, target int) bool {
	l, h := 0, len(nums)-1
	for l < h {
		var mid int = l + (h-l)/2
		if nums[mid] == target {
			return true
		}

		if nums[l] < nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				h = mid - 1
				continue
			}
			l = mid + 1
		} else if nums[mid]<nums[h] {
			if nums[mid] < target && target <= nums[h] {
				l = mid + 1
				continue
			}
			h = mid - 1
		}else {
			if nums[l]==nums[mid]{
				l++
			}
			if nums[mid]==nums[h]{
				h--
			}
		}
	}
	if nums[l]==target{
		return true
	}
	return false
}
func main() {
	nums := []int{2, 5, 6, 0, 0, 1,2}
	target :=3
	println(search(nums, target))
	nums=[]int{2,2,2,2,3,3,3,3,3,5,5,5,6,6,6,6,7,0,0,1,1,1,1,1}
	for i:=0;i<10;i++{
		println("i: ",i,
			search(nums, i))
	}
	nums=[]int{1,0,1,1,1}
	target=0
	println(search(nums, target))

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target = 13
	println(search(nums, target))

	nums=[]int{3,1}
	target=1
	println(search(nums, target))
}
