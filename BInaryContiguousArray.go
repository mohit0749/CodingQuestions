package main

import "fmt"

func main() {
	ar := []int{0, 1}
	fmt.Println(findMaxLength(ar))
}

func findMaxLength(nums []int) int {
	max := 0
	mp := make(map[int]int)
	count := 0
	for i, v := range nums {
		if v == 0 {
			count--
		} else if v == 1 {
			count++
		}
		if count == 0 {
			if max < i+1 {
				max = i + 1
			}
		}
		if pi, ok := mp[count]; ok {
			if max < i-pi {
				max = i - pi
			}
		} else {
			mp[count] = i
		}

	}
	return max
}
