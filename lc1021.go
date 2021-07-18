package main

import "fmt"

func removeOuterParentheses(s string) string {
	stack:=make([]string,0)
	var res =""
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			stack=append(stack,string(s[i]))
		}else if s[i]==')'{
			if len(stack)==1{
				res+=stack[len(stack)-1][1:]
				stack=stack[:len(stack)-1]
			}else {
				top:=stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1] += top+")"
			}
		}
	}
	return res
}


// (()())(())(()(()))

// st=[(,(,(]
//      res="()()()()"



func main(){
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}
