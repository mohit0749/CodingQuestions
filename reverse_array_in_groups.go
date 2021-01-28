package main

import "fmt"

func reverse(arr []int, start, end int) {
	if end >= len(arr) {
		end = len(arr) - 1
	}
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverseInGroups(arr []int, k int) {
	for i := 0; i < len(arr); i += k {
		reverse(arr, i, i+k-1)
	}
}

func main() {
	arr := []int{5, 6, 8, 9}
	//reverse(arr, 0, 3)
	reverseInGroups(arr, 3)
	fmt.Printf("%+v\n", arr)
}
