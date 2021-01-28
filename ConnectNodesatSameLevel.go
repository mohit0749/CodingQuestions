package main

import (
	"codes/src/queue"
)

type treeNode struct {
	val                    int
	nextRight, left, right *treeNode
}

//Using O(n) space complexity
func connectNodes(root *treeNode) {
	if root == nil {
		return
	}
	q := queue.NewQueue()
	q.Push(root)
	q.Push(nil)
	foundNil := false
	var prevNode *treeNode
	for !q.IsEmpty() {
		node, err := q.Top()
		q.Pop()
		if err != nil {
			return
		}
		tNode, ok := node.(*treeNode)
		if ok {
			foundNil = false
			if prevNode == nil {
				prevNode = tNode
			} else {
				prevNode.nextRight = tNode
				prevNode = tNode
			}
			if tNode.left != nil {
				q.Push(tNode.left)
			}
			if tNode.right != nil {
				q.Push(tNode.right)
			}
		} else if tNode == nil && !foundNil {
			foundNil = true
			prevNode = nil
			q.Push(nil)
		}
	}
}

//using O(1) constant extra space
func iterativeConnectNode(root *treeNode) {
	if root == nil {
		return
	}
	var p, q *treeNode
	p = root
	for p != nil {
		q = p
		for q != nil {
			if q.left != nil {
				if q.right != nil {
					q.left.nextRight = q.right
				} else {
					q.left.nextRight = getNextRight(q)
				}
			} else if q.right != nil {
				q.right.nextRight = getNextRight(q)
			}
			q = q.nextRight
		}
		if p.left != nil {
			p = p.left
		} else if p.right != nil {
			p = p.right
		} else {
			p = getNextRight(p)
		}
	}
}

func getNextRight(root *treeNode) *treeNode {
	node := root.nextRight
	for node != nil {
		if node.left != nil {
			return node.left
		}
		if node.right != nil {
			return node.right
		}
		node = node.nextRight
	}
	return nil
}

func main() {
	root := &treeNode{10, nil, nil, nil}
	leftSub := &treeNode{20, nil, &treeNode{40, nil, nil, nil}, &treeNode{60, nil, nil, nil}}
	right := &treeNode{30, nil, nil, nil}
	root.left = leftSub
	root.right = right
	//connectNodes(root)
	iterativeConnectNode(root)
}
