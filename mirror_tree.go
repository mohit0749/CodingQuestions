package main

import "codes/src/treebuidler"

func isMirror(t1, t2 *treebuidler.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if (t1 == nil && t2 != nil) || (t2 == nil && t1 != nil) {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func main() {
	root := treebuidler.BuildTree([]interface{}{5, 10, 10, 20, nil, 10, 20})
	println(isMirror(root.Left, root.Right))
}
