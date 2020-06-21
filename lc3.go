package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	mp := make(map[int32]int, 0)
	start := -1
	maxLen := 0
	for i, c := range s {
		if v, ok := mp[c]; ok && start < v {
			start = v
		}
		mp[c] = i
		if maxLen < i-start {
			maxLen = i - start
		}
	}
	return maxLen
}

func main() {
	s := "abcabcbbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
