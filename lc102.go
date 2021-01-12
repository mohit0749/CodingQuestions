package main

import (
	"codes/src/queue"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	q := queue.NewQueue()

	q.Push(root)
	q.Push(nil)
	level := make([]int, 0)
	var prev_nil bool
	for !q.IsEmpty() {
		item, _ := q.Top()
		_ = q.Pop()
		node, ok := item.(*TreeNode)
		if !ok {
			if prev_nil {
				break
			}
			q.Push(nil)
			prev_nil = true
			result = append(result, level)
			level = make([]int, 0)
			continue
		}
		prev_nil = false
		level = append(level, node.Val)
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
	}
	return result
}

func main() {
	//root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Printf("%+v\n", levelOrder(nil))
}
