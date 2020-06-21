package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	} else if len(cost) == 1 {
		return cost[0]
	}
	a := cost[0]
	b := cost[1]
	min := a
	if b < a {
		min = b
	}
	for i := 2; i < len(cost); i++ {
		tmpCost := cost[i] + min
		a, b = b, tmpCost
		if a < b {
			min = a
		} else {
			min = b
		}
	}
	if a < b {
		return a
	}
	return b
}

func main() {
	ar := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(ar))
}
