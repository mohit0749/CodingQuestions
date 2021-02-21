package main

type FreqStack struct {
	freq    map[int]int
	group   map[int][]int
	maxFreq int
}

func Constructor() FreqStack {
	var stk FreqStack
	stk.group = make(map[int][]int)
	stk.freq = make(map[int]int)
	return stk
}

func (stk *FreqStack) Push(x int) {
	stk.freq[x]++
	v, _ := stk.freq[x]
	if v > stk.maxFreq {
		stk.maxFreq = v
	}
	_, ok := stk.group[v]
	if !ok {
		stk.group[v] = make([]int, 0)
	}
	stk.group[v] = append(stk.group[v], x)
}

func (stk *FreqStack) Pop() int {
	list, _ := stk.group[stk.maxFreq]
	elem := list[len(list)-1]
	stk.group[stk.maxFreq] = stk.group[stk.maxFreq][:len(list)-1]
	stk.freq[elem]--
	if len(stk.group[stk.maxFreq]) == 0 {
		delete(stk.group, stk.maxFreq)
		stk.maxFreq--
	}
	return elem
}

func main() {
	obj := Constructor()
	obj.Push(4)
	obj.Push(0)
	obj.Push(9)
	obj.Push(3)
	obj.Push(4)
	obj.Push(2)
	obj.Pop()
	obj.Push(6)
	obj.Pop()
	obj.Push(1)
	obj.Pop()
	obj.Push(4)
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
}
