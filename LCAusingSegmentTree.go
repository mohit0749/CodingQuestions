package main

import (
	segmenttree "codes/src/minsegmenttree"
)

type node struct {
	data        int
	left, right *node
}

func EulerTour(root *node, ptr *int, levelPtr int, firstOcc map[int]int, level, eulerTour []int) {
	if root != nil {
		eulerTour[*ptr] = root.data
		level[*ptr] = levelPtr
		*ptr++
		if _, ok := firstOcc[root.data]; !ok {
			firstOcc[root.data] = *ptr - 1
		}
		if root.left != nil {
			EulerTour(root.left, ptr, levelPtr+1, firstOcc, level, eulerTour)
			eulerTour[*ptr] = root.data
			level[*ptr] = levelPtr
			*ptr++
		}
		if root.right != nil {
			EulerTour(root.right, ptr, levelPtr+1, firstOcc, level, eulerTour)
			eulerTour[*ptr] = root.data
			level[*ptr] = levelPtr
			*ptr++
		}
	}
}

func FindLCA(root *node, n1, n2 int) int {
	var v = 9
	var arSize = 2*v - 1
	et := make([]int, arSize, arSize)
	firstOcc := make(map[int]int)
	level := make([]int, arSize, arSize)
	var ptr = 0
	EulerTour(root, &ptr, 0, firstOcc, level, et)
	l, _ := firstOcc[n1]
	r, _ := firstOcc[n2]
	if l > r {
		l, r = r, l
	}
	st := segmenttree.NewSegmentTree(2*arSize - 1)
	st.Construct(level)
	index := st.GetMin(level, l, r)
	return et[index]
}

func main() {
	root := node{data: 1}
	leftSubtree := &node{data: 2, left: &node{data: 4}, right: &node{data: 5, left: &node{data: 8}, right: &node{data: 9}}}
	rightSubtree := &node{data: 3, left: &node{data: 6}, right: &node{data: 7}}
	root.left = leftSubtree
	root.right = rightSubtree
	println(FindLCA(&root, 3, 9))
}
