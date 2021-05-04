package main

import "sort"

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
type levelOrder struct {
	level int
	nodes []int
}

func levelOrderBottom(root *TreeNode) [][]int {
	mp := make(map[int][]int)
	preOrder(root, mp, 0)
	order := make([]levelOrder, 0)
	for k, v := range mp {
		order = append(order, levelOrder{k, v})
	}
	sort.Slice(order, func(i, j int) bool {
		if order[i].level < order[j].level {
			return false
		}
		return true
	})
	ans := make([][]int, 0)
	for _, v := range order {
		ans = append(ans, v.nodes)
	}
	return ans
}

func preOrder(root *TreeNode, mp map[int][]int, level int) {
	if root != nil {
		if _, ok := mp[level]; !ok {
			mp[level] = make([]int, 0)
		}
		mp[level] = append(mp[level], root.Val)
		preOrder(root.Left, mp, level+1)
		preOrder(root.Right, mp, level+1)
	}
}

func main() {

}
