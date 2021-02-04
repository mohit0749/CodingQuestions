package main

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minWindow(s string, t string) string {
	l := 0
	r := 0
	mp := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}

	count := 0
	ans := math.MaxInt64
	var ansL, ansR int
	tmpMp := make(map[uint8]int)
	for r < len(s) {
		if v, ok := mp[s[r]]; ok {
			tmpMp[s[r]]++
			if v1, ok := tmpMp[s[r]]; ok && v == v1 {
				count++
			}
		}
		for l<=r && count == len(mp) {
			if ans > r-l+1 {
				ans = r - l + 1
				ansL = l
				ansR = r
			}
			if v, ok := mp[s[l]]; ok {
				tmpMp[s[l]]--
				if v1, ok := tmpMp[s[l]]; ok && v1 < v {
					count--
				}
			}
			l++
		}
		r++
	}
	if ans==math.MaxInt64{
		return ""
	}
	return s[ansL:ansR+1]
}

func main() {
	s := "a"
	t := "aa"
	println(minWindow(s, t))
}
