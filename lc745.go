package main

import "fmt"

/*
Approach:

Intuition and Algorithm

Consider the word 'apple'. For each suffix of the word, we could insert that suffix, followed by '#', followed by the word, all into the trie.

For example, we will insert '#apple', 'e#apple', 'le#apple', 'ple#apple', 'pple#apple', 'apple#apple' into the trie. Then for a query like prefix = "ap", suffix = "le", we can find it by querying our trie for le#ap.
*/

type trieNode struct {
	weight int
	alpha  [27]*trieNode
}

type WordFilter struct {
	head *trieNode
}

func Constructor(words []string) WordFilter {
	var filter WordFilter
	filter.head = &trieNode{}
	for i, word := range words {
		for _, suffWord := range generateSuffixWord(word) {
			filter.addWord(suffWord, i)
		}
	}
	return filter
}

func generateSuffixWord(word string) []string {
	words := make([]string, 0)
	for i := 0; i <= len(word); i++ {
		newWord := fmt.Sprintf("%s#%s", word[len(word)-i:], word)
		words = append(words, newWord)
	}
	return words
}

func (this *WordFilter) F(prefix string, suffix string) int {
	newWord := fmt.Sprintf("%s#%s", suffix, prefix)
	head := this.head
	for _, c := range newWord {
		currWord := c - 'a'
		if c == '#' {
			currWord = 26
		}
		if head.alpha[currWord] == nil {
			return -1
		}
		head = head.alpha[currWord]
	}
	return head.weight
}

func (t *WordFilter) addWord(s string, weight int) {
	head := t.head
	for _, c := range s {
		currChar := c - 'a'
		if c == '#' {
			currChar = 26
		}
		currNode := head.alpha[currChar]
		if currNode == nil {
			head.alpha[currChar] = &trieNode{}
		}
		head = head.alpha[currChar]
		head.weight=weight
	}
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
func main() {
	//wordFilter := Constructor([]string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"})
	wordFilter := Constructor([]string{"apple"})
	//fmt.Printf("%+v\n", wordFilter.F("app", "lica"))
	//fmt.Printf("%+v\n", wordFilter.F("bccbacbcba", "a"))
	fmt.Printf("%+v\n", wordFilter.F("ap", "le"))
	//fmt.Printf("%+v\n", wordFilter.F("a", "aa"))
	//fmt.Printf("%+v\n", wordFilter.F("cabaaba","abaaaa"))
	//fmt.Printf("%+v\n", wordFilter.F("cacc", "accbbcbab"))
	//fmt.Printf("%+v\n", wordFilter.F("ccbcab", "bac"))
	//fmt.Printf("%+v\n", wordFilter.F("bac", "cba"))
	//fmt.Printf("%+v\n", wordFilter.F("ac", "accabaccaa"))
	//fmt.Printf("%+v\n", wordFilter.F("bcbb","aa"))
	//fmt.Printf("%+v\n", wordFilter.F("ccbca", "cbcababac"))

}
