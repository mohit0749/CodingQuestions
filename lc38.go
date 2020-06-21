package main

import "fmt"

func main() {
	generateSequnce()
	fmt.Println(countAndSay(4))
}

var seq [30]string

func generateSequnce() {
	seq[0] = "1"
	seq[1] = "11"
	seq[2] = "21"
	seq[3] = "1211"
	seq[4] = "111221"
	for i := 5; i < 30; i++ {
		seq[i] = countFreq(seq[i-1])
	}
}

func countFreq(s string) string {
	ss := ""
	c := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			c++
		} else {
			ts := fmt.Sprintf("%d%c", c, s[i-1])
			ss += ts
			c = 1
		}
	}
	ts := fmt.Sprintf("%d%c", c, s[len(s)-1])
	ss += ts
	return ss
}

func countAndSay(n int) string {
	return seq[n-1]
}
