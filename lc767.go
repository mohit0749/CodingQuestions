package main

import (
	"fmt"
	"sort"
)

type pair struct {
	c   byte
	cnt int
}

func reorganizeString(S string) string {
	mp := make(map[byte]int)
	ss := make([]byte, 0)
	countPair := make([]pair, 0)
	for _, c := range S {
		mp[byte(c)]++
	}
	for k := range mp {
		if mp[k] > (len(S)+1)/2 {
			return ""
		}
		countPair = append(countPair, pair{c: k, cnt: mp[k]})
	}
	sort.Slice(countPair, func(i, j int) bool {
		if countPair[i].cnt <= countPair[j].cnt {
			return true
		}
		return false
	})
	for _, pair := range countPair {
		tmp := make([]byte, 0)
		for i := 0; i < pair.cnt; i++ {
			tmp = append(tmp, pair.c)
		}
		ss = append(ss, tmp...)
	}
	var ans = make([]byte, len(S))
	var cnt int = len(S) / 2
	for i := 0; cnt < len(S); i += 2 {
		ans[i] = ss[cnt]
		cnt++

	}
	cnt = 0
	for i := 1; cnt < len(S)/2; i += 2 {
		ans[i] = ss[cnt]
		cnt++
	}
	return string(ans)
}

func main() {
	s := "ogccckcwmbmxtsbmozli"
	//s := "abbabbaaab"
	println(reorganizeString(s))
}
