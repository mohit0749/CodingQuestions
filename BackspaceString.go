package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	return getString(S) == getString(T)
}

func getString(s string) string {
	ar1 := make([]byte, len(s))
	sp := -1
	for _, c := range s {
		if c == '#' {
			if sp > -1 {
				sp--
			}
		} else {
			sp++
			ar1[sp] = byte(c)
		}
	}
	Stemp := ""
	for ; sp >= 0; sp-- {
		Stemp += string(ar1[sp])
	}
	return Stemp
}

func main() {
	s := "ab##"
	s1 := "c#d#"
	fmt.Println(backspaceCompare(s, s1))
}
