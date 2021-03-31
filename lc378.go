package main

import "container/heap"

/*
// [1,      5,     11]
// [9,     12,     13]
// [12,    13,     15]

// find 1st smallest
// after 1st smallest there are 2 possibilies (5,9) so 5 is the 2nd smallest
// to find 2nd smallest there are again two possibilies (9,11) so 9 is the 2nd smallest
// to find 3rd smallest there are these possiblities (11,12,12) so 11 is the 3rd smallest
// to find 4th smallest there are these possiblities (12,12,13) so 12 is the 4th smallest
// to find 5th smallest there are these possiblities (12,13,) so 12 is the 5th smallest
*/

// An Item is something we manage in a priority queue.
type Item struct {
	value      int
	idxI, idxJ int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func kthSmallest(matrix [][]int, k int) int {
	smallestCnt := 0
	pq := make(PriorityQueue, 0)
	mp := make(map[Item]bool)
	smallestNum := matrix[0][0]
	heap.Push(&pq, &Item{value: smallestNum, idxI: 0, idxJ: 0})
	mp[Item{smallestNum, 0, 0}] = true
	for len(pq) > 0 && smallestCnt < k {
		item := heap.Pop(&pq).(*Item)
		i, j := item.idxI, item.idxJ
		smallestNum = item.value
		smallestCnt++
		if i+1 < len(matrix) {
			item = &Item{
				value: matrix[i+1][j],
				idxI:  i + 1,
				idxJ:  j,
			}
			if _, ok := mp[*item]; !ok {
				heap.Push(&pq, item)
				mp[*item] = true
			}

		}
		if j+1 < len(matrix[i]) {
			item = &Item{
				value: matrix[i][j+1],
				idxI:  i,
				idxJ:  j + 1,
			}
			if _, ok := mp[*item]; !ok {
				heap.Push(&pq, item)
				mp[*item] = true
			}
		}
	}
	return smallestNum
}

func main() {
	matrix := [][]int{
		{1, 3, 5},
		{6, 7, 12},
		{11, 14, 14},
	}
	println(kthSmallest(matrix, 6))
}
