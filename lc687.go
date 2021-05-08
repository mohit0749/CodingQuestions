package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, mx := univaluePath(root)
	return mx - 1
}

func univaluePath(root *TreeNode) ([]int, int) {
	if root == nil {
		return []int{}, 0
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}, 1
	}
	leftPath, lmx := univaluePath(root.Left)
	rightPath, rmx := univaluePath(root.Right)

	sum := 1
	leftPathLen := len(leftPath)
	rightPathLen := len(rightPath)
	currPath := make([]int, 1)
	currPath[0] = root.Val
	var leftEqual, rightEqual bool
	if len(leftPath) > 0 && leftPath[len(leftPath)-1] == root.Val {
		sum += len(leftPath)
		leftPathLen++
		leftEqual = true
	}
	if len(rightPath) > 0 && rightPath[len(rightPath)-1] == root.Val {
		sum += len(rightPath)
		rightPathLen++
		rightEqual = true
	}
	currMax := max(leftPathLen, max(rightPathLen, sum))
	currMax = max(max(lmx, rmx), currMax)
	if leftEqual && rightEqual {
		if len(leftPath) < len(rightPath) {
			currPath = append(currPath, rightPath...)
		} else {
			currPath = append(currPath, leftPath...)
		}
	} else if leftEqual {
		currPath = append(currPath, leftPath...)
	} else if rightEqual {
		currPath = append(currPath, rightPath...)
	}

	return currPath, currMax
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:  5,
				Left: nil,
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	println(longestUnivaluePath(root))
}
