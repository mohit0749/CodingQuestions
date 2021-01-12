package main

import "math"

/**
 * Definition for a binary tree node.

 */

func isValidBST(root *TreeNode) bool {
	return isBST(root.Left, math.MinInt64, root.Val) && isBST(root.Right, root.Val, math.MaxInt64)
}

func isBST(root *TreeNode, start, end int) bool {
	if root == nil {
		return true
	}
	if root.Val > start && root.Val < end {
		leftCond, rightCond := true, true
		if root.Left != nil {
			leftCond = isBST(root.Left, start, root.Val)
		}
		if root.Right != nil {
			rightCond = isBST(root.Right, root.Val, end)
		}
		return leftCond && rightCond
	}
	return false
}

func main() {
	root := TreeNode{0, nil, nil}
	root.Left = &TreeNode{-1, nil, nil}

	println(isValidBST(&root))
}
