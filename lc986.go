package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}

	ans := make([][]int, 0)
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		p1, p2 := firstList[i], secondList[j]
		intersect := []int{max(p1[0], p2[0]), min(p1[1], p2[1])}
		if intersect[0] <= intersect[1] {
			ans = append(ans, intersect)
		}
		if p1[1] < p2[1] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func main() {
	// firstList := [][]int{{0, 2}, {4, 10}, {13, 23}, {24, 25}}
	// secondList := [][]int{{3, 5}, {8, 12}, {15, 24}, {25, 26}}
	firstList := [][]int{{5, 10}}
	secondList := [][]int{{5, 6}}
	fmt.Printf("%+v\n", intervalIntersection(firstList, secondList))
}
