package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	for head!=nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	prev, curr := head, head.Next
	for prev != nil && curr != nil {
		isFound := false
		for curr != nil && curr.Val == val {
			isFound = true
			curr = curr.Next
		}
		if isFound {
			prev.Next = curr
		}
		if curr != nil {
			curr = curr.Next
		}
		prev = prev.Next
	}
	return head
}

func printListNodes(head *ListNode) {
	for head != nil {
		print(head.Val, "->")
		head = head.Next
	}
	println()
}

func main() {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val: 5,
	//					Next: &ListNode{
	//						Val:  6,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}



	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	printListNodes(head)
	head = removeElements(head, 2)
	printListNodes(head)
}