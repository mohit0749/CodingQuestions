package main

import "codes/src/treebuidler"

func countLeaves(root *treebuidler.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return countLeaves(root.Left) + countLeaves(root.Right)
}

func main() {
	root := treebuidler.BuildTree([]interface{}{4, 8, 10, 7, nil, 5, 1, 3})
	println(countLeaves(root))
}
