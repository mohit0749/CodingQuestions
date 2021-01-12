package main

import (
	"fmt"
	"math"
)

var bitCnt map[int]int

func countBits(num int) []int {
	res := make([]int, 0)
	for i := 0; i <= num; i++ {
		//println("i:", i, " = ", checkRange(i))
		res = append(res, checkRange(i))
	}
	return res
}

func checkRange(n int) int {
	if n == 0 {
		return 0
	}
	if val, ok := bitCnt[n]; ok {
		return val
	}
	m := int(math.Log2(float64(n)))
	exactPow := int(math.Pow(2, float64(m)))
	rem := n - exactPow
	res := 1 + checkRange(rem)
	bitCnt[n] = res
	return res
}

func main() {
	arr := []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2}
	for i, v := range arr {
		fmt.Printf("i:%d = %d\n", i, v)
	}
	bitCnt = make(map[int]int)
	n := 20
	countBits(n)
	//fmt.Printf("%+v\n", )
}
