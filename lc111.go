package main

import "math"

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

func minDepth(root *TreeNode) int {
	minDepth := math.MaxInt32
	shortestDepth(root, 1, &minDepth)
	return minDepth
}

func shortestDepth(root *TreeNode, depth int, minDep *int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			if *minDep < depth {
				*minDep = depth
			}
		} else {
			shortestDepth(root.Left, depth+1, minDep)
			shortestDepth(root.Right, depth+1, minDep)
		}
	}
}

func main() {

}
