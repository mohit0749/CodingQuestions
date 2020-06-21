package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("(])"))
}

func isValid(s string) bool {
	ar := make([]rune, 0)
	top := -1
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			ar = append(ar, c)
			top++
			continue
		}
		if top < 0 {
			return false
		}
		switch c {
		case ')':
			if ar[top] == '(' {
				ar = ar[:top]
				top--
				continue
			}
		case '}':
			if ar[top] == '{' {
				ar = ar[:top]
				top--
				continue
			}
		case ']':
			if ar[top] == '[' {
				ar = ar[:top]
				top--
				continue
			}
		}
		return false
	}
	return top == -1
}
