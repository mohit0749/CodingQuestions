package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(arr []string) string {
	sort.Slice(arr, func(i, j int) bool {
		s1 := arr[i] + arr[j]
		s2 := arr[j] + arr[i]
		n1, _ := strconv.ParseInt(s1, 10, 32)
		n2, _ := strconv.ParseInt(s2, 10, 32)
		if n1 > n2 {
			return true
		}
		return false
	})
	return strings.Join(arr, "")
}

func main() {
	//arr := []string{"3", "30", "34", "5", "9"}
	arr := []string{"54", "546", "548", "60"}
	println(largestNumber(arr))
}
