package main

import "fmt"

/*
["((()))","(()())","(())()","()(())","()()()"]
*/

func generateParenthesis(n int) []string {
	return printParenthesis(n, 0, 0, "")
}

func printParenthesis(n int, opened, closed int, currStr string) []string {
	if n == 0 && opened == closed {
		return []string{currStr}
	}
	parens := make([]string, 0)
	if n == 0 && closed < opened {
		parens = append(parens, printParenthesis(n, opened, closed+1, currStr+")")...)
	} else if n > 0 {
		if opened > closed {
			parens = append(parens, printParenthesis(n-1, opened+1, closed, currStr+"(")...)
			parens = append(parens, printParenthesis(n, opened, closed+1, currStr+")")...)
		} else if opened == closed {
			parens = append(parens, printParenthesis(n-1, opened+1, closed, currStr+"(")...)
		}
	}
	return parens
}

func main() {
	fmt.Printf("%+v\n", generateParenthesis(3))
}
