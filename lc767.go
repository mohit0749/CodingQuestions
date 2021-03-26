package main

import (
	"sort"
	"strings"
)

func reorganizeString(S string) string {
	s:=[]byte(S)
	sort.Slice(s,func(i,j int)bool{
		if s[i]!=s[j]{return true}
		return false
	})
	var ss strings.Builder
	for _,c:=range s{
		ss.WriteRune(rune(c))
	}
	return ss.String()
}

func main(){

}
