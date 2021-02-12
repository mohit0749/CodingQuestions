package main

import (
	"codes/src/treebuidler"
	"math"
)

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

func preoder(root *treebuidler.TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	preoder(root.Left)
	preoder(root.Right)
}
func main() {
	arr := []interface{}{20, 8, 22, 5, 3, nil, 25, nil, nil, 10, 14}
	root := treebuidler.BuildTree(arr)
	preoder(root)
}
