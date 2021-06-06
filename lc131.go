package main

import "fmt"

func partition(s string) [][]string {
	results := make([][]string, 0)
	results = backtrack(s, make([]string, 0), &results, make(map[string][][]string))
	return results
}

func backtrack(s string, part []string, results *[][]string, mp map[string][][]string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	palin := ""
	partition := make([][]string, 0)
	for i, v := range s {
		palin += string(v)
		if v, ok := mp[s]; ok {
			partition = v
		} else if isPalindrom(palin) {
			partition = append(partition, backtrack(s[i+1:], append([]string{}, append(part, palin)...), results, mp))
		}
	}
	for i := 0; i < len(partition); i++ {

	}
	return partition
}

func isPalindrom(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
