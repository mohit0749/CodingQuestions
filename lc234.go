package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	var prev *ListNode
	for {
		//handler odd number case
		if fast == nil {
			fast = slow.Next
			slow = prev
			break
		}

		//handle even number case
		if fast.Next == nil {
			fast = slow.Next
			slow.Next = prev
			break
		}
		fast = fast.Next.Next
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp

	}
	for fast != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//1->2->3->1->nil
func main() {
	head := &ListNode{
		Val: 1,
		Next: nil,
	}
	println(isPalindrome(head))
}
