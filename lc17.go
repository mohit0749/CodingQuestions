package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	return getCombinations(digits, 0, "", digitMap)
}

func getCombinations(digits string, start int, combStr string, digitMap map[byte]string) []string {
	if start >= len(digits) {
		if combStr == "" {
			return []string{}
		}
		return []string{combStr}
	}
	comb := make([]string, 0)
	for _, c := range digitMap[digits[start]] {
		comb = append(comb, getCombinations(digits, start+1, combStr+string(c), digitMap)...)
	}
	return comb
}

func main() {
	digits := "23"
	fmt.Printf("%+v\n", letterCombinations(digits))
}
