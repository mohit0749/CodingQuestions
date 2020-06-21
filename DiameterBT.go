package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var MaxDia int

func diameterOfBinaryTree(root *TreeNode) int {
	MaxDia = 0
	maxDiaMeter(root)
	return MaxDia
}

func maxDiaMeter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	if root.Left != nil {
		left = maxDiaMeter(root.Left) + 1
	}
	if root.Right != nil {
		right = maxDiaMeter(root.Right) + 1
	}
	MaxDia = max(MaxDia, left+right)
	return max(left, right)

}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(diameterOfBinaryTree(root))
}
