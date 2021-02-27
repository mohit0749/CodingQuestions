package main

import (
	"fmt"
	"strings"
	"unicode"
)

func letterCasePermutation(S string) []string {

	return getPermutation(S, 0)
}

func flipCase(c rune) rune {
	if unicode.IsLower(c) {
		return unicode.ToUpper(c)
	}
	return unicode.ToLower(c)
}

func getPermutation(s string, i int) []string {
	if i >= len(s) {
		return []string{s}
	}
	permuts := make([]string, 0)
	permuts = append(permuts, getPermutation(s, i+1)...)
	if !unicode.IsDigit(rune(s[i])) {
		var newS strings.Builder
		newS.WriteString(s[:i])
		newS.WriteRune(flipCase(rune(s[i])))
		newS.WriteString(s[i+1:])
		permuts = append(permuts, getPermutation(newS.String(), i+1)...)
	}
	return permuts
}

func main() {
	S := "a1b2"
	fmt.Printf("%+v\n", letterCasePermutation(S))

}
