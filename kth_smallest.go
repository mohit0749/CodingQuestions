package main

import (
	"fmt"
	"math/rand"
)

func quickSelect(arr []int, k, l, h int) int {
	if l == h {
		return arr[l]
	} else if l > h {
		return -1
	}
	pivot := partition(arr, l, h)
	if pivot == k-1 {
		return arr[pivot]
	}
	if k-1 < pivot {
		return quickSelect(arr, k, l, pivot-1)
	} else {
		return quickSelect(arr, k, pivot+1, h)
	}
}

func partition(arr []int, l, h int) int {
	randIndex := rand.Intn(h - l + 1)
	arr[h], arr[l+randIndex] = arr[l+randIndex], arr[h]
	i := l
	for j := l; j <= h; j++ {
		if arr[j] < arr[h] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[h] = arr[h], arr[i]
	return i
}

func main() {
	arr := []int{7, 10, 4, 20, 15}
	fmt.Printf("%+v\n", quickSelect(arr, 4, 0, len(arr)-1))
}
