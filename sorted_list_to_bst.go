package main

type ListNode struct {
	data int
	next *ListNode
}
func NewListNode(data int, next *ListNode)*ListNode{
	return &ListNode{data,next}
}

func printList(head *ListNode){
	for head!=nil{
		print(head.data,"->")
	}
}


type TreeNode struct {
	data  int
	left, right *TreeNode
}

func NewTreeNode(data int,left,right *TreeNode)*TreeNode{
	return &TreeNode{data,left,right}
}

func traverseTree(root *TreeNode){
	if root==nil{
		return
	}
	traverseTree(root.left)
	print(root.data," ")
	traverseTree(root.right)
}

//O(n+n) approach

func countNodes(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head=head.next
	}
	return n
}

func constructBST(head *ListNode, n int)*TreeNode{
	if n<=0{
		return nil
	}
	leftSTree:=constructBST(head,n/2)
	root:=NewTreeNode(head.data,leftSTree,nil)
	*head=*head.next
	root.right=constructBST(head,n-n/2-1)
	return root
}


func main() {
	head:=&ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,&ListNode{5,&ListNode{6,&ListNode{7,nil}}}}}}}
	printList(head)
	n:=countNodes(head)
	root:=constructBST(head,n)
	traverseTree(root)
}
