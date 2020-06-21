package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	startPtr := head
	nextPtr := head
	for startPtr != nil {
		nextPtr = startPtr
		for {
			if nextPtr.Next != nil && nextPtr.Next.Val == startPtr.Val {
				nextPtr = nextPtr.Next
			} else {
				break
			}
		}
		startPtr.Next = nextPtr.Next
		startPtr = startPtr.Next
	}
	return head
}
func main() {
	node := ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}}}
	node = *deleteDuplicates(&node)
	for {
		fmt.Println(node.Val)
		if node.Next == nil {
			break
		}
		node = *node.Next
	}
}
