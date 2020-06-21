package main

import "fmt"

type StackNode struct {
	val        int
	prev, next *StackNode
}

type MinStack struct {
	minEle int
	Stk    *StackNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if this.Stk == nil {
		this.minEle = x

		this.Stk = &StackNode{
			val:  x,
			prev: nil,
			next: nil,
		}
		return
	}
	if x <= this.minEle {
		tmp := x
		x = 2*x - this.minEle
		this.minEle = tmp
	}
	this.Stk.next = &StackNode{
		val:  x,
		prev: this.Stk,
		next: nil,
	}
	this.Stk = this.Stk.next
}

func (this *MinStack) Pop() {
	if this.Stk == nil {
		return
	}
	if this.Stk.val < this.minEle {
		this.minEle = 2*this.minEle - this.Stk.val
	}
	this.Stk = this.Stk.prev
	if this.Stk != nil {
		this.Stk.next = nil
	}
}

func (this *MinStack) Top() int {
	if this.Stk.val < this.minEle {
		return this.minEle
	}
	return this.Stk.val
}

func (this *MinStack) GetMin() int {
	return this.minEle
}

func main() {
	minStack := Constructor()
	//fmt.Println(minStack.GetMin())
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	//minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	//fmt.Println(minStack.GetMin())
	minStack.Top()
	fmt.Println(minStack.GetMin())
	//minStack.Pop()
	//fmt.Println(minStack.GetMin())
}
