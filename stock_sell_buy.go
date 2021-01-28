package main

import "fmt"

func buyAndSell(price []int) [][]int {
	if len(price) == 0 || len(price) == 1 {
		return [][]int{}
	}
	pairs := make([][]int, 0)
	minVal := 0
	minIdx := 0
	if price[0] < price[1] {
		minVal = price[0]
		minIdx = 0
	} else {
		minVal = price[1]
		minIdx = 1
	}
	for i := minIdx + 1; i < len(price); i++ {
		if price[i-1] < price[i] {
			if minVal > price[i-1] {
				minVal = price[i-1]
				minIdx = i - 1
			}
		} else if price[i-1] > price[i] {
			if minIdx != i-1 && price[i-1] > minVal {
				pairs = append(pairs, []int{minIdx, i - 1})
				minVal = price[i]
				minIdx = i
			}

		}
	}
	if minIdx != len(price)-1 && price[len(price)-1] > minVal {
		pairs = append(pairs, []int{minIdx, len(price) - 1})
	}
	return pairs
}

func main() {
	//arr := []int{7,1,5,3,6,4}
	//arr := []int{1,2,3,4,5}
	//arr := []int{7,6,4,3,1}
	//arr := []int{100,180,260,310,40,535,695}
	//arr := []int{4,2,2,2,4}
	arr := []int{2}
	fmt.Printf("%+v\n", buyAndSell(arr))
}
