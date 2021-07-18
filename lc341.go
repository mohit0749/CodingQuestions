package main

import "fmt"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	n    int
	list []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	if this.list == nil {
		return true
	}
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return this.n
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.n = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	if this.list == nil {
		this.list = make([]*NestedInteger, 0)
	}
	this.list = append(this.list, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	if this.list == nil {
		return make([]*NestedInteger, 0)
	}
	return this.list
}

type Queue struct {
	arr []int
}

func NewQueue() *Queue {
	return &Queue{make([]int, 0)}
}

func (q *Queue) Push(x int) {
	q.arr = append(q.arr, x)
}

func (q *Queue) Pop() int {
	ele := -1
	if !q.IsEmpty() {
		ele = q.arr[0]
		q.arr = q.arr[1:]
	}
	return ele
}

func (q *Queue) IsEmpty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) String() string {
	return fmt.Sprintf("%+v\n", q.arr)
}

type NestedIterator struct {
	q    *Queue
	it   int
	list []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{NewQueue(), 0, nestedList}
}

func (this *NestedIterator) Next() int {
	if !this.q.IsEmpty() {
		return this.q.Pop()
	}
	fmt.Println("in Next")
	this.getAllInteger(this.list[this.it])
	fmt.Println(this.q)
	this.it++
	return this.q.Pop()
}

func (this *NestedIterator) getAllInteger(list *NestedInteger) {
	if len(list.GetList()) == 0 {
		this.q.Push(list.GetInteger())
		return
	}
	for i := 0; i < len(list.GetList()); i++ {
		this.getAllInteger(list.GetList()[i])
	}

}

func (this *NestedIterator) HasNext() bool {
	if this.q.Size() > 0 || this.it < len(this.list) {
		return true
	}
	return false
}

func main() {
	var list = make([]*NestedInteger, 0)

	for i := 0; i < 3; i++ {
		var n NestedInteger
		list = append(list, &n)
	}
	list[0].Add(NestedInteger{1, nil})
	list[0].Add(NestedInteger{1, nil})
	list[1].SetInteger(2)
	list[2].Add(NestedInteger{1, nil})
	list[2].Add(NestedInteger{1, nil})
	res := make([]int, 0)
	it := Constructor(list)
	for it.HasNext() {
		res = append(res, it.Next())
	}
	fmt.Printf("%+v\n", res)
}
