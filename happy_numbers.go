package main

import "fmt"

func isHappy(n int) bool {
	var mp = make(map[int]bool)
	var sqrt = squareIt(n)
	for !mp[sqrt] {
		mp[sqrt] = true
		if sqrt == 1 {
			return true
		}
		sqrt = squareIt(sqrt)
	}
	return false
}
func squareIt(n int) int {
	f := genDigit(&n)
	i := f()
	sqrt := i * i
	for n > 0 {
		i = f()
		sqrt += i * i
	}
	fmt.Println(sqrt)
	return sqrt
}
func genDigit(n *int) func() int {
	return func() int {
		i := *n % 10
		*n /= 10
		return i
	}
}

func main() {
	fmt.Println(isHappy(900))
}
