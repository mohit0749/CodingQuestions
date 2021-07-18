package main

import "fmt"

func scoreOfParentheses(s string) int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	for i, _ := range s {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else {
			stk := stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*stk, 1)
		}
	}
	return stack[len(stack)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//"()"
	//"(())"
	//"()()"
	//"(()(()))"

	//s:="()"
	//s:="(())"
	//s := "()()"
	s := "(()(()))"
	fmt.Println(scoreOfParentheses(s))
}
