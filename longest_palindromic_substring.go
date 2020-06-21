package main

import "fmt"

/*
Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

func main() {
	s := "babad"
	s = "cbbd"
	s = "forgeeksskeegfor"
	s = "Geeks"
	s = "ccc"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	} else if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return ""
	}
	dp := make([][]int, len(s))
	maxLen := 1
	var ss string = s[0:1]
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
		if i+1 < len(s) {
			if i+1 < len(s) && s[i] == s[i+1] {
				dp[i][i+1] = 1
				maxLen = 2
				ss = s[i : i+maxLen]
			}
		}
	}
	for j := 2; j < len(s); j++ {
		for i := 0; i < len(s)-j; i++ {
			if i+j < len(s) && s[i] == s[i+j] && dp[i+1][i+j-1] == 1 {
				dp[i][i+j] = 1
				if maxLen < j+1 {
					maxLen = j + 1
					ss = s[i : i+j+1]
				}
			}
		}
	}
	return ss
}

func printarr(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Print(dp[i][j], ",")
		}
		fmt.Print("\n")
	}
}
