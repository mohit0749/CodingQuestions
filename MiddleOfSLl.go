package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}
func NewSLL(ar ...int) *ListNode {
	var head, node *ListNode
	node = NewListNode(0)
	head = node
	for _, i := range ar {
		node.Next = NewListNode(i)
		node = node.Next
	}
	return head.Next
}

func traverse(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	ar := []int{1, 2, 3}
	head := NewSLL(ar...)
	//traverse(head)
	head = middleNode(head)
	traverse(head)
}

func middleNode(head *ListNode) *ListNode {
	var h1, h2 *ListNode
	h1, h2 = head, head
	for h2 != nil {
		h2 = h2.Next
		if h2 != nil {
			h2 = h2.Next
			h1 = h1.Next
		}
	}
	return h1
}
