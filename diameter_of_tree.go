package main

import "codes/src/treebuidler"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func diameter(root *treebuidler.TreeNode, dist *int) int {
	if root == nil {
		return 0
	}
	left := diameter(root.Left, dist)
	right := diameter(root.Right, dist)
	*dist = max(*dist, 1+left+right)
	return 1 + max(left, right)
}

func getDiameter(root *treebuidler.TreeNode) int {
	dist := 0
	_ = diameter(root, &dist)
	return dist
}

func main() {
	root := treebuidler.BuildTree([]interface{}{10, 20, 30})
	println(getDiameter(root))
}
