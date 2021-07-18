package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Stack struct {
	arr []*TreeNode
}

func NewStack() Stack {
	return Stack{make([]*TreeNode, 0)}
}

func (stk *Stack) Push(x *TreeNode) {
	stk.arr = append(stk.arr, x)
}

func (stk *Stack) Pop() *TreeNode {
	if len(stk.arr) == 0 {
		return nil
	}
	ele := stk.arr[len(stk.arr)-1]
	stk.arr = stk.arr[:len(stk.arr)-1]
	return ele
}

func (stk *Stack) Top() *TreeNode {
	if stk.Size() >= 1 {
		return stk.arr[stk.Size()-1]
	}
	return nil
}

func (stk *Stack) Size() int {
	return len(stk.arr)
}

func (stk *Stack) IsEmpty() bool {
	return len(stk.arr) == 0
}

func (stk Stack) String() string {
	fmt.Printf("%+v\n", stk.arr)
	return ""
}

type BSTIterator struct {
	stk Stack
}

func Constructor(root *TreeNode) BSTIterator {
	stack := NewStack()
	stack.Push(root)
	root = root.Left
	for root != nil {
		stack.Push(root)
		root = root.Left
	}
	return BSTIterator{stack}
}

func (this *BSTIterator) Next() int {
	//fmt.Println(this.stk.Size(),this.stk)
	ele := this.stk.Pop()
	if ele != nil && ele.Right != nil {
		this.stk.Push(ele.Right)
		for this.stk.Top() != nil && this.stk.Top().Left != nil {
			this.stk.Push(this.stk.Top().Left)
		}
	}

	return ele.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.stk.IsEmpty()
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
	}
	obj := Constructor(root)
	println(obj.Next())
	println(obj.Next())
	println(obj.HasNext())
	println(obj.Next())
	println(obj.HasNext())
	println(obj.Next())
	println(obj.HasNext())
	println(obj.Next())
	println(obj.HasNext())
}

///["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]
//[[[7,3,15,null,null,9,20]],[],[],[],[],[],[],[],[],[]]
