package main

import "fmt"

//without using any extra space using O(m*n) time complexity and Insertion sort like approach
func merge(arr1, arr2 []int) {
	i := 0
	j := 0
	for i < len(arr1) {
		if arr1[i] > arr2[j] {
			arr1[i], arr2[j] = arr2[j], arr1[i] //swapping
			// then sort 2nd array using insertion sort
			insertionSort(arr2, j)
		}
		i++
	}
}

func insertionSort(arr []int, i int) {
	j := i + 1
	tmp := arr[i]
	for tmp > arr[j] {
		arr[j-1] = arr[j]
		j++

	}
	arr[j-1] = tmp
}

//using O(m+n) approach use shell sort algorithm
func shellSort(arr1, arr2 []int) {
	n1 := len(arr1)
	n2 := len(arr2)
	var gap int = (n1 + n2) / 2
	var i int
	for ; gap > 0; gap /= 2 {
		for ; i+gap < len(arr1); i++ {
			if arr1[i+gap] < arr1[i] {
				arr1[i+gap], arr1[i] = arr1[i], arr1[i+gap]
			}
		}

		for ; i < len(arr1); i++ {
			if i+gap > len(arr1) {
				j := (i + gap) - len(arr1) - 1
				if j > len(arr2) {
					break
				}
				if arr1[i] > arr2[j] {
					arr1[i], arr2[j] = arr2[j], arr1[i]
				}
			}
		}

		for ; i+gap < len(arr2); i++ {
			if arr2[i+gap] < arr2[i] {
				arr2[i+gap], arr2[i] = arr2[i], arr2[i+gap]
			}
		}
	}
}

func main() {
	// arr1 := []int{1, 3, 5, 7}
	// arr2 := []int{0, 2, 6, 8, 9}
	// insertionSort(arr, 2)
	arr1 := []int{10, 12}
	arr2 := []int{5, 8, 20}
	// merge(arr1, arr2)
	shellSort(arr1, arr2)
	fmt.Printf("%+v, %+v\n", arr1, arr2)
}
