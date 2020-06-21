package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var mbuy = prices[0]
	var mprofit = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < mbuy {
			mbuy = prices[i]
			continue
		}
		var profit = prices[i] - mbuy
		if profit > mprofit {
			mprofit = profit
		}
	}
	return mprofit
}

func main() {
	ar := []int{5, 1, 3, 2, 1, 0, 7}
	fmt.Println(maxProfit(ar))
}
