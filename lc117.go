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
			if root.Right != nil {
				root.Left.Next = root.Right
			} else {
				root.Left.Next = nextRight(root.Next)
			}
		}
		if root.Right != nil {
			root.Right.Next = nextRight(root.Next)
		}
		if root.Right != nil {
			connectLevel(root.Right)
		}

		if root.Left != nil {
			connectLevel(root.Left)
		}
	}
}

func nextRight(root *Node) *Node {
	for root != nil {
		if root.Left != nil {
			return root.Left
		} else if root.Right != nil {
			return root.Right
		}
		root = root.Next
	}
	return nil
}

func main() {
	root := &Node{Val: 0,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 1,
				Left: &Node{
					Val:   5,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Right: &Node{
					Val:   1,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Right: nil,
			Next:  nil,
		},
		Right: &Node{
			Val: 4,
			Left: &Node{
				Val:  3,
				Left: nil,
				Right: &Node{
					Val:   6,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Right: &Node{
				Val:  -1,
				Left: nil,
				Right: &Node{
					Val:   8,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Next: nil,
		},
	}
	connect(root)
}
