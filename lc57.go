package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 && len(newInterval) == 0 {
		return [][]int{}
	} else if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if newInterval[0] > intervals[len(intervals)-1][1] {
		intervals = append(intervals, newInterval)
	}
	ans := make([][]int, 0)
	start := newInterval[0]
	end := newInterval[1]
	posFound := false
	isOverlapped := false
	pos := 0
	for _, v := range intervals {
		pos++
		if posFound {
			ans = append(ans, v)
			continue
		}
		if v[0] >= start && v[1] <= end {
			isOverlapped = true
			continue
		}
		if start >= v[0] && end <= v[1] {
			posFound = true
			isOverlapped = false
			ans = append(ans, v)
		} else if start < v[0] && end < v[0] {
			posFound = true
			isOverlapped = false
			ans = append(ans, []int{start, end})
			ans = append(ans, v)
		} else if start <= v[0] && end >= v[0] && end <= v[1] {
			end = v[1]
			isOverlapped = true
		} else if start >= v[0] && start <= v[1] && end > v[1] {
			isOverlapped = true
			start = v[0]
		} else {
			ans = append(ans, v)
		}
	}
	if pos == len(intervals) && isOverlapped {
		ans = append(ans, []int{start, end})
	}

	return ans
}

func main() {
	//intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	//newInterval := []int{4, 8}
	//fmt.Printf("%+v\n", insert(intervals, newInterval))
	//
	//intervals = [][]int{{1, 3}, {6, 9}}
	//newInterval = []int{2, 5}
	//fmt.Printf("%+v\n", insert(intervals, newInterval))
	//
	//intervals = [][]int{}
	//newInterval = []int{5, 7}
	//fmt.Printf("%+v\n", insert(intervals, newInterval))
	//
	//intervals = [][]int{{1, 5}}
	//newInterval = []int{2, 3}
	//fmt.Printf("%+v\n", insert(intervals, newInterval))
	//
	//intervals = [][]int{{1, 5}}
	//newInterval = []int{2, 7}
	//fmt.Printf("%+v\n", insert(intervals, newInterval))

	//intervals := [][]int{{0, 5}, {8, 9}}
	//newInterval := []int{0, 5}
	//fmt.Printf("%+v\n", insert(intervals, newInterval))

	intervals := [][]int{{1, 3},{6, 8},{9, 9}}
	newInterval := []int{7, 8}
	fmt.Printf("%+v\n", insert(intervals, newInterval))
}
