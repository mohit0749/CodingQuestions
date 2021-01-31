package main

import "codes/src/treebuidler"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func height(root *treebuidler.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}
func main() {
	root := treebuidler.BuildTree([]interface{}{1, 2, 3, 4, 5, 6, 7, 1, 2, nil, 4, 5, nil, 7, nil})
	println(height(root))
}
