package main

type llNode struct {
	val  int
	next *llNode
}

func mergeRecursive(head1, head2 *llNode) *llNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	if head1.val < head2.val {
		head1.next = mergeRecursive(head1.next, head2)
		return head1
	} else {
		head2.next = mergeRecursive(head1, head2.next)
		return head2
	}
}

func traverseLL(head *llNode) {
	for head != nil {
		println(head.val)
		head = head.next
	}
}

func main() {
	l1 := &llNode{4, &llNode{7, &llNode{9, nil}}}
	l2 := &llNode{5, &llNode{6, &llNode{8, nil}}}
	if l1.val < l2.val {
		mergeRecursive(l1, l2)
		traverseLL(l1)
	} else {
		mergeRecursive(l1, l2)
		traverseLL(l2)
	}

	//
}
