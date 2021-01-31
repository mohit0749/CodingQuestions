package main

import (
	"codes/src/treebuidler"
	"math"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxPathSum(root *treebuidler.TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := maxPathSum(root.Left, maxSum)
	right := maxPathSum(root.Right, maxSum)
	currSum := max(left+root.Val, max(right+root.Val, root.Val))
	*maxSum = max(max(*maxSum, left+root.Val+right), currSum)
	return currSum
}

func main() {
	root := treebuidler.BuildTree([]interface{}{3, 4, 5, -10, 4, nil, nil})
	//root := treebuidler.BuildTree([]interface{}{-10, 9, 20, nil, nil, 15, 7})
	//root := treebuidler.BuildTree([]interface{}{1, -2, -3, 1, 3, -2, nil, -1})
	pathSum := math.MinInt32
	maxPathSum(root, &pathSum)
	println(pathSum)
}
