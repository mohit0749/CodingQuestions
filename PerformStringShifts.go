package main

import "fmt"

func stringShift(s string, shift [][]int) string {
	rs := 0
	ls := 0
	for _, v := range shift {
		if v[0] == 1 {
			rs += v[1]
		} else if v[0] == 0 {
			ls += v[1]
		}
	}
	rs -= ls
	if rs < 0 {
		s = leftShift(s, -1*rs)
	} else {
		s = rightShift(s, rs)
	}
	return s
}
func rightShift(s string, n int) string {
	n %= len(s)
	return s[len(s)-n:] + s[:len(s)-n]
}

func leftShift(s string, n int) string {
	n %= len(s)
	return s[n:] + s[:n]
}
func main() {
	//s:="yisxjwry"
	//shift:=[][]int{{1,8},{1,4},{1,3},{1,6},{0,6},{1,4},{0,2},{0,1}}
	s := "abc"
	shift := [][]int{{0, 1}, {1, 2}}
	fmt.Print(stringShift(s, shift))
}
