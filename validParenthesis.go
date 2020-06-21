package main

import (
	"fmt"
)

func main() {
	s := "(((******))"
	fmt.Print(checkValidString(s))
}
func checkValidString(s string) bool {

	rc, lc := 0, 0
	for _, c := range s {
		if c == '(' {
			rc++
			lc++
		} else if c == ')' {
			rc--
			lc--
		} else if c == '*' {
			lc++
			rc--
		}
		if rc < 0 {
			rc = 0
		}
		if lc < 0 {
			return false
		}
	}
	if rc == 0 {
		return true
	}
	return false
}
