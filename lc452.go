package main

import (
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		}
		if points[i][0]==points[j][0] && points[i][1]<points[j][1]{
			return true
		}
		return false
	})
	start := points[0][0]
	end := points[0][1]
	cnt := 1
	for i := 1; i < len(points); i++ {
		a, b := points[i][0], points[i][1]
		if start <= a && a <= end {
			start = max(start, a)
			end=min(end,b)
			continue
		} else {
			start = a
			end = b
			cnt++
		}
	}
	return cnt
}

func main() {
	//arr := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	//println(findMinArrowShots(arr))
	//arr = [][]int{{1,2},{3,4},{5,6},{7,8}}
	//println(findMinArrowShots(arr))
	//arr = [][]int{{1,2},{2,3},{3,4},{4,5}}
	//println(findMinArrowShots(arr))
	//arr = [][]int{{1,2}}
	//println(findMinArrowShots(arr))
	//arr = [][]int{{2,3},{2,3}}
	//println(findMinArrowShots(arr))
	//
	//arr = [][]int{{-1,1},{0,1},{2,3},{1,2}}
	//println(findMinArrowShots(arr))

	//arr:=[][]int{{3,9},{7,12},{3,8},{6,8},{9,10},{2,9},{0,9},{3,9},{0,6},{2,8}}
	//println(findMinArrowShots(arr))

	//arr := [][]int{{1, 9}, {7, 16}, {2, 5}, {7, 12}, {9, 11}, {2, 10}, {9, 16}, {3, 9}, {1, 3}}
	//println(findMinArrowShots(arr))

	//arr := [][]int{{1,2},{4,5},{1,5}}
	//println(findMinArrowShots(arr))

	arr := [][]int{{9,12},{1,10},{4,11},{8,12},{3,9},{6,9},{6,7}}
	println(findMinArrowShots(arr))
}
