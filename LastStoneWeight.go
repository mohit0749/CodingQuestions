package main

import (
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Build() {
	for i := len(*h) / 2; i >= 1; i-- {
		h.heapify(i)
	}
}

func (h *IntHeap) heapify(i int) {
	largest := i
	if v, leftIndex := h.getLeftChild(i); v != -1 && v > (*h)[largest] {
		largest = leftIndex
	}
	if v, rightIndex := h.getRightChild(i); v != -1 && v > (*h)[largest] {
		largest = rightIndex
	}
	if largest != i {
		h.Swap(i, largest)
		h.heapify(largest)
	}
}

func (h IntHeap) parent(child int) (int, int) {
	return h[child/2], child / 2
}
func (h IntHeap) getLeftChild(root int) (int, int) {
	lc := 2 * root
	if lc < len(h) {
		return h[lc], lc
	}
	return -1, -1
}
func (h IntHeap) getRightChild(root int) (int, int) {
	rc := 2*root + 1
	if rc < len(h) {
		return h[rc], rc
	}
	return -1, -1
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) PopMax() int {
	old := *h
	n := len(old)
	x := old[1]
	old[1] = old[n-1]
	*h = old[:n-1]
	h.heapify(1)
	return x
}

func (h *IntHeap) Insert(val int) {
	*h = append(*h, val)
	i := len(*h) - 1
	for i > 1 {
		v, pidx := h.parent(i)
		if v < (*h)[i] {
			h.Swap(i, pidx)
			i = pidx
		} else {
			break
		}
	}
}

func lastStoneWeight(stones []int) int {
	lenSt := len(stones)
	heapStones := make(IntHeap, 1)
	heapStones = append(heapStones, stones...)
	heapStones.Build()
	destroyed := 0
	for len(heapStones) > 2 && destroyed != lenSt {
		x := heapStones.PopMax()
		y := heapStones.PopMax()
		if x == y {
			destroyed += 2
			continue
		}
		rem := 0
		if x < y {
			rem = y - x
		} else {
			rem = x - y
		}
		destroyed++
		heapStones.Insert(rem)
	}
	if destroyed == lenSt {
		return 0
	}
	return heapStones[1]
}

func main() {
	stones := []int{1, 1}
	//stones.Build()
	fmt.Println(lastStoneWeight(stones))
	//for len(stones) > 1 {
	//	fmt.Println(stones.PopMax())
	//}
}
