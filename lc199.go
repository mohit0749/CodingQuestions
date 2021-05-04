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

func rightSideView(root *TreeNode) []int {
	levels := make([]int, 1000)
	for i := 0; i < len(levels); i++ {
		levels[i] = -1
	}
	traversal(root, 0, levels)
	ans := make([]int, 0)
	for i := 0; i < len(levels); i++ {
		if levels[i] != -1 {
			ans = append(ans, levels[i])
		}
	}
	return ans
}

func traversal(root *TreeNode, level int, mp []int) {
	if root != nil {
		mp[level-1] = root.Val
		if root.Left != nil {
			traversal(root.Left, level+1, mp)
		}
		if root.Right != nil {
			traversal(root.Right, level+1, mp)
		}
	}
}
