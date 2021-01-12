package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return reverseLL(head, nil)
}

func reverseLL(head *ListNode, prevNode *ListNode) *ListNode {
	if head == nil {
		return prevNode
	}
	p := reverseLL(head.Next, head)
	head.Next = prevNode
	return p
}

func reverseIterLL(head *ListNode) *ListNode {
	var next, prev *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func print(head *ListNode) {
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	print(head)
	//newHead := reverseList(head)
	newHead := reverseIterLL(head)
	print(newHead)
}
