package main

import "fmt"

func main() {
	ar := []int{3, 2, 2, 3}
	l := removeElement(ar, 3)
	fmt.Println(ar, l)
}
func swap(a, b int) (int, int) {
	return b, a
}
func removeElement(nums []int, val int) int {

	i, j := 0, len(nums)-1
	for i <= j {
		if nums[j] == val {
			j--
			continue
		}
		if nums[i] == val {
			nums[i], nums[j] = swap(nums[i], nums[j])
			j--
		}
		i++
	}

	return i
}
