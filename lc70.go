package main

import "fmt"

func main() {
	n := 44
	fmt.Println(climbStairs(n))

}

func climbStairs(n int) int {
	climbedSteps := make([]int, 3)
	climbedSteps[0] = 0
	climbedSteps[1] = 1
	climbedSteps[2] = 2
	for i := 3; i <= n; i++ {
		climbedSteps = append(climbedSteps, climbedSteps[i-1]+climbedSteps[i-2])
	}
	return climbedSteps[n]
}
