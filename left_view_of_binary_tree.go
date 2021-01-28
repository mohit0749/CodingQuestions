package main

import "math"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func leftView(root *TreeNode, height int, maxHeight *int) {
	if root == nil {
		return
	}
	if height > *maxHeight {
		*maxHeight = height
		println(root.val)
	}
	leftView(root.left, height+1, maxHeight)
	leftView(root.right, height+1, maxHeight)
}

func main() {
	//root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, &TreeNode{8, nil, nil}}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{4, nil, nil}}}
	root := &TreeNode{10, &TreeNode{20, &TreeNode{40, nil, nil}, &TreeNode{60, nil, nil}}, &TreeNode{30, nil, nil}}
	h := math.MinInt32
	leftView(root, 0, &h)
}
