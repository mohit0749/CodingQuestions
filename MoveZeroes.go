package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	i, j := 0, 0
	for i < len(nums) && j < len(nums) {
		for i < len(nums) && nums[i] != 0 {
			i++
		}
		if i > len(nums) {
			break
		}
		if j > i && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}

}

func main() {
	ar := []int{0, 0, 1, 1, 1, 1, 12, 2}
	moveZeroes(ar)
	fmt.Println(ar)
}
