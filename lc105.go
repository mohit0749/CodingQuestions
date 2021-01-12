package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootPtr := 0
	return constructTree(preorder, inorder, &rootPtr)
}

func constructTree(preorder, inorder []int, ptr *int) *TreeNode {
	if len(inorder)==0{
		return nil
	}
	if len(inorder) == 1 {
		*ptr = *ptr + 1
		return &TreeNode{inorder[0], nil, nil}
	}
	idx := findRoot(preorder[*ptr], inorder)
	if idx == -1 {
		return nil
	}
	*ptr = *ptr + 1
	root := &TreeNode{inorder[idx], nil, nil}
	root.Left = constructTree(preorder, inorder[:idx], ptr)
	root.Right = constructTree(preorder, inorder[idx+1:], ptr)
	return root
}

func findRoot(val int, inorder []int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}

func main() {
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	root := buildTree(preorder, inorder)
	print(root)
}
