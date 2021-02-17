package main

type Trie struct {
	head *trieNode
}

type trieNode struct {
	terminal bool
	charset  [26]*trieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{head: &trieNode{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	head := this.head
	for _, c := range word {
		currChar := c - 'a'
		if head.charset[currChar] == nil {
			head.charset[currChar] = &trieNode{}
		}
		head = head.charset[currChar]
	}
	head.terminal = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	head := this.head
	for _, c := range word {
		currChar := c - 'a'
		if head.charset[currChar] == nil {
			return false
		}
		head = head.charset[currChar]
	}
	return head.terminal
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	head := this.head
	for _, c := range prefix {
		currChar := c - 'a'
		if head.charset[currChar] == nil {
			return false
		}
		head = head.charset[currChar]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	//trie.Insert("apple")
	//trie.Insert("app")
	//println(trie.Search("apple"))
	//println(trie.Search("app"))
	//trie.Insert("mohit")
	//trie.Insert("manish")
	//println(trie.Search("ppp"))
	//println(trie.StartsWith("mohit"))
	//println(trie.StartsWith(""))
	//println(trie.StartsWith("app"))
	trie.Insert("k")
	println(trie.Search("k"))

	//trie.Insert("apple")
	//println(trie.Search("apple"))
	//println(trie.Search("app"))
	//println(trie.StartsWith("app"))
	//trie.Insert("app")
	//println(trie.Search("app"))
}
