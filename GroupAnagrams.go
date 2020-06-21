package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		tmpStrSlice := strings.Split(str, "")
		sort.Strings(tmpStrSlice)
		tmpStr := strings.Join(tmpStrSlice, "")
		mp[tmpStr] = append(mp[tmpStr], str)
	}
	grpStrs := make([][]string, 0)
	for _, v := range mp {
		grpStrs = append(grpStrs, v)
	}
	return grpStrs
}

func main() {
	strs := []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(strs))
}
