package main

/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/
const MAX = int(1e4 + 1)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	mp := make(map[int]int)
	res := checkDemons(coins, amount, mp)
	if res == MAX {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkDemons(coins []int, amount int, mp map[int]int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if val, ok := mp[amount]; ok {
		return val
	}
	minCoin := MAX
	for i := len(coins) - 1; i >= 0; i-- {
		changed := checkDemons(coins, amount-coins[i], mp)
		if changed >= 0 && changed < minCoin {
			minCoin = 1 + changed
		}
	}
	mp[amount] = minCoin
	return mp[amount]
}
func main() {
	//coins := []int{2, 5, 10, 1}
	//coins := []int{1,2,5}
	coins := []int{2}
	amount := 3
	println(coinChange(coins, amount))
}
