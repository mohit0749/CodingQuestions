package main

import "math"

func kadane(arr []int) int {
	sum := 0
	maxSum := math.MinInt32
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum < 0 {
			sum = 0
			if maxSum < arr[i] {

				maxSum = arr[i]
			}
		} else if maxSum < sum {
			maxSum = sum
		}

	}
	return maxSum
}

func missingNumber(arr []int, n int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	var triangle int = (n * (n + 1)) / 2
	return triangle - sum
}

func main() {
	// arr := []int{-1, -2, -3, -4}
	// arr := []int{1, 2, 3, -2, 5}
	// arr := []int{1, 2, 3, 5}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	// println(kadane(arr))
	println(missingNumber(arr, 10))
}
