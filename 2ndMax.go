package main

import (
	"math/big"
)

func secondMax(nums []string) string {
	if len(nums) == 0 || len(nums) == 1 {
		return "-1"
	}

	first := new(big.Int)
	var second *big.Int
	first.SetString(nums[0], 10)
	for i := 1; i < len(nums); i++ {
		nextNum := new(big.Int)
		nextNum.SetString(nums[i], 10)
		if first.Cmp(nextNum) == 0 {
			continue
		}
		if first.Cmp(nextNum) == -1 {
			second = first
			first = nextNum
		} else if second == nil {
			second = nextNum
		} else if second.Cmp(nextNum) == -1 {
			second = nextNum
		}
	}
	if second == nil {
		return "-1"
	}
	return second.String()
}

func main() {
	println(secondMax([]string{"3", "-2"}))
	println(secondMax([]string{"5", "5", "4", "4", "2"}))
	println(secondMax([]string{"4", "4", "4"}))
	println(secondMax([]string{"-214748364801", "-214748364802"}))
}
