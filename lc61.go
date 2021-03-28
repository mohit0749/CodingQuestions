package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

/*
1->2->3->4->5->
k=2

1->2->

*/
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	prevHead := head
	tmpHead := head
	head = moveHead(head, k)
	//var prev *ListNode
	for head.Next != nil {
		head = head.Next
		//prev = tmpHead
		tmpHead = tmpHead.Next
	}
	//prev.Next = nil
	head.Next = prevHead
	newHead := tmpHead.Next
	tmpHead.Next = nil
	return newHead

}

func moveHead(head *ListNode, k int) *ListNode {
	cnt := 0
	newHead := head
	for i := 0; i < k; i++ {
		cnt++
		head = head.Next
		if head == nil {
			head = newHead
			return moveHead(head, k%cnt)
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
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	head.print()
	k := 7
	//moveHead(head, k).print()
	rotateRight(head, k).print()
}
