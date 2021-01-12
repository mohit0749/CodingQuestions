package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	First  int
	Second int
	Sum    int
}
type heapTuple []*Item

func (h heapTuple) Len() int {
	return len(h)
}

func (h heapTuple) Less(i, j int) bool {
	return h[i].Sum < h[j].Sum
}

func (h heapTuple) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapTuple) Pop() interface{} {
	old := *h
	n := old.Len()
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}

func (h *heapTuple) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func kSmallestPairsHeap(nums1 []int, nums2 []int, k int) [][]int {
	pairs := make([][]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return pairs
	}
	k = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}(len(nums2)*len(nums1), k)

	item := &Item{
		First:  0,
		Second: 0,
		Sum:    nums1[0] + nums2[0],
	}
	pq := heapTuple{item}
	mp := make(map[Item]bool)
	mp[*item] = true
	heap.Init(&pq)
	for k > 0 {
		item = heap.Pop(&pq).(*Item)
		i := item.First
		j := item.Second
		pairs = append(pairs, []int{nums1[i], nums2[j]})
		if j+1 < len(nums2) {
			item = &Item{
				First:  i,
				Second: j + 1,
				Sum:    nums1[i] + nums2[j+1],
			}
			if !mp[*item] {
				heap.Push(&pq, item)
				heap.Fix(&pq, pq.Len()-1)
				mp[*item] = true
			}
		}
		if i+1 < len(nums1) {
			item = &Item{
				First:  i + 1,
				Second: j,
				Sum:    nums1[i+1] + nums2[j],
			}
			if !mp[*item] {
				heap.Push(&pq, item)
				heap.Fix(&pq, pq.Len()-1)
				mp[*item] = true
			}
		}
		k--
	}
	return pairs
}

func main() {
	ar1 := []int{1, 1, 2}
	ar2 := []int{}
	k := 10
	pairs := kSmallestPairsHeap(ar1, ar2, k)
	for i, _ := range pairs {
		fmt.Println(pairs[i][0], pairs[i][1])
	}
}
