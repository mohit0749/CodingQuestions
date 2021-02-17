package main

type WordDictionary struct {
	head *trieNode
}

type trieNode struct {
	terminal bool
	charset  [27]*trieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{head: &trieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	head := this.head
	for _, c := range word {
		currChar := c - 'a'
		if c == '.' {
			currChar = 26
		}
		if head.charset[currChar] == nil {
			head.charset[currChar] = &trieNode{}
		}
		head = head.charset[currChar]
	}
	head.terminal = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.head,0,word)
}

func dfs(head *trieNode, i int, word string) bool {
	if i >= len(word) {
		if head!=nil && head.terminal==true{
			return true
		}
		return false
	}
	if head == nil && i < len(word) {
		return false
	}

	if word[i] == '.' {
		isMatched := false
		for _, node := range head.charset {
			if node != nil {
				isMatched = isMatched || dfs(node, i+1, word)
			}
		}
		return isMatched
	}

	if head.charset[word[i]-'a'] != nil {
		return dfs(head.charset[word[i]-'a'], i+1, word)
	}
	return false
}

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.AddWord("mohit")
	wordDictionary.AddWord("manish")
	wordDictionary.AddWord("apple")
	//wordDictionary.AddWord("app")

	println(wordDictionary.Search("pad")) // return False
	println(wordDictionary.Search("bad")) // return True
	println(wordDictionary.Search(".ad")) // return True
	println(wordDictionary.Search("b..")) // return True
	println(dfs(wordDictionary.head, 0, "f.d"))
}
