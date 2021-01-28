package main

type treeNode struct {
	val         int
	left, right *treeNode
}

func isBST(root *treeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min == nil && max == nil {
		return isBST(root.left, min, &root.val) && isBST(root.right, &root.val, max)
	}
	if min == nil && *max > root.val {
		return isBST(root.left, min, &root.val) && isBST(root.right, &root.val, max)
	} else if max == nil && *min < root.val {
		return isBST(root.left, min, &root.val) && isBST(root.right, &root.val, max)
	} else if *min < root.val && *max > root.val {
		return isBST(root.left, min, &root.val) && isBST(root.right, &root.val, max)
	}
	return false
}

func main() {
	//root := &treeNode{2, &treeNode{1, nil, nil}, &treeNode{3, nil, nil}}
	root := &treeNode{2, nil, &treeNode{7, nil, &treeNode{6, nil, &treeNode{5, nil, &treeNode{9, nil, nil}}}}}
	println(isBST(root, nil, nil))
}
