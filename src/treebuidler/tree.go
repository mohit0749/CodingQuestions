package treebuidler

import "codes/src/queue"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func NewTreeNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{val, left, right}
}

func BuildTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	val, ok := arr[0].(int)
	if !ok {
		return nil
	}
	root := NewTreeNode(val, nil, nil)
	rootMain := root
	q := queue.NewQueue()
	q.Push(struct {
		i    int
		root *TreeNode
	}{0, root})
	i := 0
	for !q.IsEmpty() {
		item, err := q.Top()
		if err != nil {
			return nil
		}
		_ = q.Pop()
		node := item.(struct {
			i    int
			root *TreeNode
		})
		i = node.i
		root = node.root
		left := 2*i + 1
		right := 2*i + 2
		if left < len(arr) {
			val, ok := arr[left].(int)
			if ok {
				root.Left = NewTreeNode(val, nil, nil)
				q.Push(struct {
					i    int
					root *TreeNode
				}{left, root.Left})
			}

		}
		if right < len(arr) {
			val, ok := arr[right].(int)
			if ok {
				root.Right = NewTreeNode(val, nil, nil)
				q.Push(struct {
					i    int
					root *TreeNode
				}{right, root.Right})
			}
		}
	}

	return rootMain
}
