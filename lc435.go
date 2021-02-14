package main

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	removed := 0
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] < intervals[j][1] {
			return true
		}
		if intervals[i][1] == intervals[j][1] && intervals[i][0] > intervals[j][0] {
			return true
		}
		return false
	})

	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0]<end {
			removed++
		} else {
			end = intervals[i][1]
		}
	}
	return removed
}

func main() {
	//arr := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3},{3,7}}
	//arr:=[][]int{{1,2},{2,3},{3,4},{1,3}}
	//arr := [][]int{{1, 2}, {1, 2}, {1, 2}}
	//arr := [][]int{{1, 2}, {1, 2}}
	//arr := [][]int{{1, 2}, {1, 4}}
	//arr := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}

	arr := [][]int{{4, 6}, {3, 7}, {15, 18}, {16, 17}}
	println(eraseOverlapIntervals(arr))
}
