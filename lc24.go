package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, node1, node2, newHead *ListNode

	node1, node2 = head, head.Next
	newHead = node2
	for node1 != nil && node2 != nil {
		tmpNode1 := node2.Next
		var tmpNode2 *ListNode
		if tmpNode1 != nil {
			tmpNode2 = tmpNode1.Next
		}
		swapNodes(prev, node1, node2)
		prev = node1
		node1 = tmpNode1
		node2 = tmpNode2
	}
	return newHead
}

func swapNodes(prev, node1, node2 *ListNode) {
	tmp := node2.Next
	node2.Next = node1
	if prev != nil {
		prev.Next = node2
	}
	node1.Next = tmp
}

func printListNodes(head *ListNode) {
	for head != nil {
		print(head.Val, "->")
		head = head.Next
	}
	println()
}

func main() {
	head := &ListNode{1,
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	newHead := swapPairs(head)
	printListNodes(newHead)
}
