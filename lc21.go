package main

import "fmt"

func main() {
	var l1 ListNode
	l1.Val = 0
	l1.Next = nil
	newList := mergeTwoLists(&l1, nil)
	for newList != nil {
		fmt.Print(newList.Val, "->")
		newList = newList.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ptr1, ptr2 *ListNode
	ptr1 = l1
	ptr2 = l2
	var newList ListNode
	p := &newList
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val >= ptr2.Val {
			p.Val = ptr2.Val
			p.Next = &ListNode{}
			ptr2 = ptr2.Next
			p = p.Next
		} else if ptr1.Val < ptr2.Val {
			p.Val = ptr1.Val
			p.Next = &ListNode{}
			ptr1 = ptr1.Next
			p = p.Next
		}
	}
	if ptr1 == nil {
		p.Next = ptr2
	} else if ptr2 == nil {
		p.Val = ptr1.Val
		p.Next = ptr1.Next
	}
	return &newList
}
