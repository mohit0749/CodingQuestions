package main

import "fmt"

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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructTree(nums)
}

func constructTree(nums []int) *TreeNode {
	idx := findMax(nums)
	if idx != -1 {
		fmt.Println(nums[idx])
		root := &TreeNode{nums[idx], nil, nil}
		root.Left = constructTree(nums[:idx])
		root.Right = constructTree(nums[idx+1:])
		return root
	}
	return nil
}

func findMax(nums []int) int {
	max := -1
	idx := -1
	for i, v := range nums {
		if max < v {
			max = v
			idx = i
		}
	}
	return idx
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	constructMaximumBinaryTree(nums)
}
