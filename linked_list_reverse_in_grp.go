package main

type llNode struct {
	val  int
	next *llNode
}

func reverseLL(head *llNode, newHead, newTail, next **llNode, k int) *llNode {
	if k == 1 || head.next == nil {
		*newHead = head
		*next = head.next
		head.next = nil
		return head
	}
	prevNode := reverseLL(head.next, newHead, newTail, next, k-1)
	prevNode.next = head
	*newTail = head
	(*newTail).next = nil
	return head
}

func reverseInGroup(head *llNode, k int) {

	var next *llNode
	var tmpHead *llNode
	var currHead = head
	var prevTail *llNode
	for {
		var newHead, newTail *llNode
		reverseLL(currHead, &newHead, &newTail, &next, k)
		if tmpHead == nil {
			tmpHead = newHead
		}
		if prevTail==nil{
			prevTail=newTail
		}else{
			prevTail.next=newHead
		}
		currHead.next = next
		currHead = next
		if next == nil {
			break
		}
	}
	traverseLL(tmpHead)
}

func traverseLL(head *llNode) {
	for head != nil {
		println(head.val)
		head = head.next
	}
}

func main() {
	head := &llNode{1, &llNode{2, &llNode{3, &llNode{4, &llNode{5,nil}}}}}
	reverseInGroup(head, 3)
}
