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

func height(root *treebuidler.TreeNode, isBalanced *bool) int {
	if root == nil {
		return 0
	}
	left := height(root.Left, isBalanced)
	right := height(root.Right, isBalanced)
	diff := float64(left - right)
	if math.Abs(diff) > 1 {
		*isBalanced = false
	}
	return 1 + max(left, right)
}

func isBalanced(root *treebuidler.TreeNode) bool {
	balanced := true
	height(root, &balanced)
	return balanced
}

func main() {
	root := treebuidler.BuildTree([]interface{}{1, 10, nil, 5})
	println(isBalanced(root))
}
