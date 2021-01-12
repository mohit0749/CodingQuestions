package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	maxDepth := 0
	depth(root, 0, &maxDepth)
	return maxDepth
}

func depth(root *TreeNode, currDepth int, maxDepth *int) {
	if root == nil {
		if *maxDepth < currDepth {
			*maxDepth = currDepth
		}
		return
	}
	depth(root.Left, currDepth+1, maxDepth)
	depth(root.Right, currDepth+1, maxDepth)
}

func main() {
	//root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	root := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	print(maxDepth(root))
}
