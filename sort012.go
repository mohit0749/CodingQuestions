package main

import "fmt"

func sort012(arr []int) {
	i := 0
	l := 0
	k := len(arr) - 1
	for l <= k {
		v := arr[l]
		if v == 0 {
			arr[l], arr[i] = arr[i], arr[l]
			i++
			l++
		} else if v == 2 {
			arr[l], arr[k] = arr[k], arr[l]
			k--
		} else if v == 1 {
			l++
		}
	}
}

func main() {
	arr := []int{0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 0, 0, 0, 1, 1, 1}
	sort012(arr)
	fmt.Printf("%+v\n", arr)
}
