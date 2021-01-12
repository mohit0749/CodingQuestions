package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return check(p, q)
}

func check(rootP *TreeNode, rootQ *TreeNode) bool {
	if rootP == nil && rootQ == nil {
		return true
	} else if rootP == nil && rootQ != nil {
		return false
	} else if rootP != nil && rootQ == nil {
		return false
	}

	if rootP.Val == rootQ.Val {
		return check(rootP.Left, rootQ.Left) && check(rootP.Right, rootQ.Right)
	}

	return false
}

func main() {
	root1 := &TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{1, nil, nil}}
	root2 := &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
	println(isSameTree(root1, root2))
}
