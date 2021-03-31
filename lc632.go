package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value      int
	idxI, idxJ int
}

type MaxHeap []*Item

func (pq *MaxHeap) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MaxHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func (pq MaxHeap) Len() int {
	return len(pq)
}

func (pq MaxHeap) Less(i, j int) bool {
	return pq[i].value > pq[j].value
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type MinHeap []*Item

func (pq MinHeap) Len() int {
	return len(pq)
}

func (pq MinHeap) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinHeap) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func smallestRange(nums [][]int) []int {
	var minPq = make(MinHeap, 0)
	var maxPq = make(MaxHeap, 0)
	for i := 0; i < len(nums); i++ {
		item := &Item{nums[i][0], i, 0}
		heap.Push(&minPq, item)
		heap.Push(&maxPq, item)
	}
	var smallestRange []int
	for len(minPq) > 0 && len(maxPq) > 0 {
		minItem := heap.Pop(&minPq).(*Item)
		maxItem := maxPq[0]
		tmpI := minItem.value
		tmpJ := maxItem.value
		if smallestRange == nil {
			smallestRange = make([]int, 2)
			smallestRange[0] = tmpI
			smallestRange[1] = tmpJ
		} else if tmpJ-tmpI < smallestRange[1]-smallestRange[0] {
			smallestRange[0] = tmpI
			smallestRange[1] = tmpJ
		} else if tmpJ-tmpI == smallestRange[1]-smallestRange[0] {
			if tmpI < smallestRange[0] {
				smallestRange[0] = tmpI
				smallestRange[1] = tmpJ
			}
		}
		if minItem.idxI == maxItem.idxJ && maxItem.idxJ == maxItem.idxJ && minItem.value == maxItem.value {
			break
		}
		if minItem.idxJ+1 < len(nums[minItem.idxI]) {
			item := &Item{
				value: nums[minItem.idxI][minItem.idxJ+1],
				idxI:  minItem.idxI,
				idxJ:  minItem.idxJ + 1,
			}
			heap.Push(&minPq, item)
			heap.Push(&maxPq, item)
		} else {
			break
		}
	}
	return smallestRange
}

func main() {
	//nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	//nums := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	//nums := [][]int{{10},{11}}
	nums := [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}}
	fmt.Printf("%+v\n", smallestRange(nums))
}
