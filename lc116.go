package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	connectLevel(root)
	return root
}

func connectLevel(root *Node) {
	if root != nil {
		if root.Left != nil {
			root.Left.Next = root.Right
			if root.Next != nil {
				root.Right.Next = root.Next.Left
			}
		}
		if root.Left != nil {
			connectLevel(root.Left)
		}
		if root.Right != nil {
			connectLevel(root.Right)
		}
	}
}
