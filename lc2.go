package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//n1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	//n2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	n1 := ListNode{1, nil}
	n2 := ListNode{1, &ListNode{1, &ListNode{1, nil}}}
	newList := addTwoNumbers(&n1, &n2)
	for {
		fmt.Println(newList.Val)
		if newList.Next == nil {
			break
		}
		newList = newList.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	head := res
	var carry int
	for {
		currentDigit := l1.Val + l2.Val + carry
		carry = int(currentDigit / 10)
		currentDigit %= 10
		res.Val = currentDigit
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil || l2 == nil {
			break
		}
		res.Next = &ListNode{0, nil}
		res = res.Next
	}
	tmpNode := l1
	res.Next = l1
	if l2 != nil {
		res.Next = l2
		tmpNode = l2
	}
	for carry > 0 && tmpNode != nil {
		tmpNode.Val += carry
		carry = int(tmpNode.Val / 10)
		tmpNode.Val %= 10
		res = res.Next
		if tmpNode.Next == nil {
			break
		}
		tmpNode = tmpNode.Next

	}
	if carry > 0 {
		res.Next = &ListNode{carry, nil}
	}

	return head
}
