package main

import "log"

func main() {
	var s, t string
	s = "abc"
	t = "ahbgdc"
	log.Print(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	var i, j int
	i = 0
	j = 0
	for ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}
