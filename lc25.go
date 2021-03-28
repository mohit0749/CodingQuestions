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

func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead, prev, newNext *ListNode
	start := head
	end := head
	for {
		cnt := 1
		end = start
		for end != nil {
			end = end.Next
			cnt++
			if cnt == k {
				break
			}
		}
		if end == nil {
			break
		}
		if newHead == nil {
			newHead = end
		}
		newNext = end.Next
		reverse(start, end)
		if prev != nil {
			prev.Next = end
		}
		prev = start
		prev.Next = newNext
		start = newNext
	}
	if newHead == nil {
		newHead = head
	}
	return newHead
}

//1->2->3->4->5->nil reverse to
// 5->4->3->2->1->nil
func reverse(start, end *ListNode) *ListNode {
	if start == end {
		return start
	}
	prev := reverse(start.Next, end)
	prev.Next = start
	start.Next = nil
	return start
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
	var newHead *ListNode
	newHead = reverseKGroup(head, 3)
	newHead.print()
}
