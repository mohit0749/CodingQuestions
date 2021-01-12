package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	pathSum(root, 0, &maxSum)
	return maxSum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func pathSum(root *TreeNode, currMax int, maxSum *int) int {
	if root == nil {
		return 0
	}
	leftSum := pathSum(root.Left, currMax, maxSum)
	rightSum := pathSum(root.Right, currMax, maxSum)
	currMax = max(leftSum+root.Val, max(root.Val, rightSum+root.Val))
	*maxSum = max(*maxSum, max(max(root.Val+leftSum, root.Val+rightSum), max(root.Val, root.Val+leftSum+rightSum)))
	return currMax
}
func main() {
	//
	root := &TreeNode{1,
		&TreeNode{-2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	println(maxPathSum(root))
}
