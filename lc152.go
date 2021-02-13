package main

func findMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxProd := nums[0]
	posProd := nums[0]
	negProd := nums[0]
	for i := 1; i < len(nums); i++ {
		if posProd*nums[i] > posProd {
			posProd *= nums[i]
		} else {
			posProd = nums[i]
		}
		negProd *= nums[i]
		if posProd > maxProd {
			maxProd = posProd
		}
		if negProd > maxProd {
			maxProd = negProd
		} else if negProd == 0 {
			negProd = nums[i]
		}
	}
	if posProd > maxProd {
		maxProd = posProd
	}
	if negProd > maxProd {
		maxProd = negProd
	}
	return maxProd
}

func maxProduct(nums []int) int {
	reverse := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		reverse = append(reverse, nums[i])
	}
	maxProd1 := findMax(nums)
	maxProd2 := findMax(reverse)
	if maxProd1 < maxProd2 {
		return maxProd2
	}
	return maxProd1
}

func main() {
	nums := []int{-3, -2, 2, -5}
	println(maxProduct(nums))
}
