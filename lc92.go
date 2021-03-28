package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) print() {
	for h != nil {
		print(h.Val, "->")
		h = h.Next
	}
	println()
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var prev, next *ListNode
	reverse(head, &prev, &next, left, right, 1)
	if left == 1 {
		return next
	}
	return head
}

func reverse(head *ListNode, leftPrev, rightNext **ListNode, left, right, count int) *ListNode {
	if head == nil {
		return nil
	}
	if count == right {
		*rightNext = head
		(*leftPrev).Next = head.Next
		return head
	}
	if count == left {
		*leftPrev = head
	}
	prev := reverse(head.Next, leftPrev, rightNext, left, right, count+1)
	if count+1 == left {
		head.Next = *rightNext
	}
	if count >= left && count <= right && *leftPrev != nil && *rightNext != nil {
		if prev != nil {
			prev.Next = head
		}
	}
	return head
}

func main() {
	head := &ListNode{
		1,
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
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	head.print()
	newHead := reverseBetween(head, 1, 2)
	newHead.print()
}
