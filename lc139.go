package main

func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]bool)
	for _, w := range wordDict {
		mp[w] = true
	}
	que := make([]int, 0) //use as a queue
	que = append(que, 0)
	visited:=make([]bool,len(s),len(s))
	for len(que) > 0 {
		startIdx := que[0]
		que = que[1:]
		if visited[startIdx]{
			continue
		}
		for i := startIdx + 1; i <= len(s); i++ {
			segment := s[startIdx:i]
			if _, ok := mp[segment]; ok {
				que = append(que, i)
				if i == len(s) {
					return true
				}
			}
		}
		visited[startIdx]=true

	}
	return false
}
func main() {
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	s := "catsandog"
	println(wordBreak(s, wordDict))
}
