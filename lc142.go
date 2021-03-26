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

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	slow = head.Next
	fast = head.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {
	//3,2,0,-4
	head := ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	head.Next.Next.Next.Next = head.Next
	cycleNode := detectCycle(&head)
	println(cycleNode.Val)
}
