package main

import (
	"container/heap"
	//"math"

	//"fmt"
	//"strconv"

	//"fmt"
	//"math"
	//"strconv"
)

type Segment struct {
	start, end int
	index      int
	maxDist    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Segment

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].maxDist == pq[j].maxDist {
		return pq[i].start < pq[j].end
	}
	return pq[i].maxDist > pq[j].maxDist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Segment)
	item.index = n
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

type ExamRoom struct {
	seats int
	pq    PriorityQueue
}

func Constructor(n int) ExamRoom {
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Segment{-1, n, 0, n})
	return ExamRoom{n, pq}
}

func (er *ExamRoom) getDist(start, end int) int {
	if start == 1 {
		return end
	} else if end == er.seats {
		return end - start - 1
	}
	return (end - start) / 2
}

func (er *ExamRoom) Seat() int {
	if len(er.pq) == 0 {
		heap.Push(&er.pq, &Segment{0, 2 * (er.seats - 1), 0, er.seats})
		return 0
	}
	var item = heap.Pop(&er.pq).(*Segment)
	seatNo := -1
	seatNo = item.start + (item.end-item.start)/2
	if seatNo != 0 {
		heap.Push(&er.pq, &Segment{item.start, seatNo, 0, er.getDist(item.start, seatNo)})
	}
	if seatNo != er.seats-1 {
		heap.Push(&er.pq, &Segment{seatNo, item.end, 0, er.getDist(seatNo, item.end)})
	}
	return seatNo
}

func (er *ExamRoom) Leave(p int) {
	//(dist,x1,y1) (d,x2,y2) => (d+d,x1,y2) --> merge two overlapping interval
	var head, tail *Segment
	for i := 0; i < len(er.pq); i++ {
		item := er.pq[i]
		if item.end == p {
			head = item
		}
		if item.start == p {
			tail = item
		}
		if head != nil && tail != nil {
			break
		}
	}
	heap.Remove(&er.pq, head.index)
	heap.Remove(&er.pq, tail.index)
	heap.Push(&er.pq, &Segment{
		start:   head.start,
		end:     tail.end,
		index:   0,
		maxDist: er.getDist(head.start, tail.end),
	})
}

func main() {
	obj := Constructor(4)
	//["ExamRoom","seat","seat","seat","seat","leave","leave","seat"]
	//	[[4],[],[],[],[],[1],[3],[]]
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Leave(1)
	obj.Leave(3)
	obj.Seat()

	//obj := Constructor(10)
	//["ExamRoom","seat","seat","leave","seat","seat","seat","seat","seat","seat","seat","seat","seat","leave","seat","leave","seat","leave","seat","leave","seat","leave","seat","leave","leave","seat","seat","leave","seat","leave","seat","leave","seat","leave","leave","seat"]
	//[[10],[],[],[9],[],[],[],[],[],[],[],[],[],[9],[],[1],[],[8],[],[2],[],[1],[],[0],[1],[],[],[7],[],[3],[],[7],[],[0],[3],[]]
	//res := make([]string, 0)
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(9)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(9)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(1)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(8)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(2)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(1)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(0)
	//res = append(res, "null")
	//obj.Leave(1)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(7)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(3)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(7)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//obj.Leave(0)
	//res = append(res, "null")
	//obj.Leave(3)
	//res = append(res, "null")
	//res = append(res, strconv.Itoa(obj.Seat()))
	//fmt.Println(res)

	//["ExamRoom","seat","seat","seat","leave","leave","seat","seat","seat","seat","seat","seat","seat"]
	//	[[8],[],[],[],[0],[7],[],[],[],[],[],[],[]]
	//obj := Constructor(8)
	//obj.Seat()
	//obj.Seat()
	//obj.Seat()
	//obj.Leave(0)
	//obj.Leave(7)
	//obj.Seat()
	//obj.Seat()
	//obj.Seat()
	//obj.Seat()
	//obj.Seat()
	//obj.Seat()
	//obj.Seat()
}

//["ExamRoom","seat","seat","seat","leave","leave","seat","seat","seat","seat","seat","seat","seat","seat","seat","leave"]
//[[10],[],[],[],[0],[4],[],[],[],[],[],[],[],[],[],[0]]
/**
 * Your ExamRoom object will be instantiated and called as such:
[null,0,9,null,9,4,2,6,1,3,5,7,8,null,9,null,1,null,8,null,2,null,1,null,null,0,1,null,7,null,3,null,7,null,null,0]
 *
 * param_1 := res=append(res,strconv.Itoa(obj.Seat()));
 * obj.Leave(p);
*/
