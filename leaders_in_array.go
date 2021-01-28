package main

import "fmt"

func printLeaders(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	currMax := arr[len(arr)-1]
	leaders := make([]int, 0)
	leaders = append(leaders, currMax)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < currMax {
			continue
		}
		currMax = max(currMax, arr[i])
		leaders = append(leaders, arr[i])
	}
	return leaders
}

func main() {
	arr := []int{1, 2, 3, 4, 0}
	fmt.Println("%+v\n", printLeaders(arr))
}
