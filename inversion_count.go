package main

import "fmt"

func merge(arr []int, l, mid, h int) int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for i := l; i <= mid; i++ {
		arr1 = append(arr1, arr[i])
	}
	for i := mid + 1; i <= h; i++ {
		arr2 = append(arr2, arr[i])
	}
	fmt.Printf("%+v,%+v", arr1, arr2)
	i := 0
	j := 0
	k := l
	inv := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			k++
			i++
		} else if arr2[j] < arr1[i] {
			inv += len(arr1) - i
			arr[k] = arr2[j]
			j++
			k++
		}
	}

	for i < len(arr1) {
		arr[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		arr[k] = arr2[j]
		j++
		k++
	}

	println("= ", inv)
	return inv
}

func countInversion(arr []int, l, h int) int {
	var inv int
	if l >= h {
		return 0
	}
	var mid int = (l + h) / 2
	inv += countInversion(arr, l, mid)
	inv += countInversion(arr, mid+1, h)
	inv += merge(arr, l, mid, h)
	return inv
}

func main() {
	// arr := []int{2, 3, 5, 1, 6}
	arr := []int{2, 4, 1, 3, 5}
	// arr := []int{2, 3, 4, 5, 6}
	// inv := merge(arr, 0, 2, 5)
	inv := countInversion(arr, 0, len(arr)-1)
	fmt.Printf("%+v, %+v", inv, arr)
}
