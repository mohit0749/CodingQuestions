package minsegmenttree

import "fmt"

//SegmentTree
type SegmentTree struct {
	st []int
}

func NewSegmentTree(size int) *SegmentTree {
	var sTree SegmentTree
	sTree.st = make([]int, size, size)
	return &sTree
}

func (t *SegmentTree) Construct(arr []int) {
	t.buildTree(arr, 0, 0, len(arr)-1)
}

func (t *SegmentTree) buildTree(arr []int, stIdx, startIdx, endIdx int) {
	if startIdx == endIdx {
		t.st[stIdx] = startIdx
		return
	}
	var mid int = (startIdx + endIdx) / 2
	leftChildIdx, rightChildIdx := stIdx*2+1, stIdx*2+2
	t.buildTree(arr, leftChildIdx, startIdx, mid)
	t.buildTree(arr, rightChildIdx, mid+1, endIdx)
	if arr[t.st[leftChildIdx]] < arr[t.st[rightChildIdx]] {
		t.st[stIdx] = t.st[leftChildIdx]
	} else {
		t.st[stIdx] = t.st[rightChildIdx]
	}
}

func (t *SegmentTree) GetMin(arr []int, l, r int) int {
	return t.min(arr, l, r, 0, len(arr)-1, 0)
}

func (t *SegmentTree) min(arr []int, l, r, startIdx, endIdx, stIdx int) int {
	if l <= startIdx && r >= endIdx {
		return t.st[stIdx]
	}
	if l > endIdx || r < startIdx {
		return -1
	}
	var mid int = (startIdx + endIdx) / 2
	leftMin := t.min(arr, l, r, startIdx, mid, stIdx*2+1)
	rightMin := t.min(arr, l, r, mid+1, endIdx, stIdx*2+2)
	if leftMin == -1 {
		return rightMin
	} else if rightMin == -1 {
		return leftMin
	}
	if arr[leftMin] < arr[rightMin] {
		return leftMin
	}
	return rightMin
}

func (t SegmentTree) printTree() {
	fmt.Printf("%+v\n", t.st)
}

func (t *SegmentTree) Update(arr []int, value, at int) {
	t.update(arr, 0, len(arr)-1, 0, value, at)
}

func (t *SegmentTree) update(arr []int, startIdx, endIdx, stIdx, value, at int) {
	if at < startIdx || at > endIdx {
		return
	}
	if t.st[stIdx] > value {
		t.st[stIdx] = value
	}
	if startIdx != endIdx {
		var mid int = (startIdx + endIdx) / 2
		t.update(arr, startIdx, mid, stIdx*2+1, value, at)
		t.update(arr, mid+1, endIdx, stIdx*2+2, value, at)
	}
}
