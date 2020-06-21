package main

import "fmt"

func main() {
	n := 8
	fmt.Println(mySqrt(n))
}

func mySqrt(x int) int {
	for i := 1; i <= x; i++ {
		sqrt := i * i
		if sqrt == x {
			return i
		} else if sqrt > x {
			return i - 1
		}
	}
	return 1
}
