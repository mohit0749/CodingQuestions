package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	t := make([]int,0)
	inorderTraversal(root, &t)
	return t[k-1]
}

func inorderTraversal(root *TreeNode, traversal *[]int) {
	if root != nil {
		inorderTraversal(root.Left, traversal)
		*traversal = append(*traversal, root.Val)
		inorderTraversal(root.Right, traversal)
	}
}

func main() {
	root := &TreeNode{
		3,
		&TreeNode{1, nil, &TreeNode{4, nil, nil}},
		&TreeNode{4, nil, nil},
	}
	print(kthSmallest(root, 1))
}
