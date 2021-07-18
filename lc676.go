package main

type MagicDictionary struct {
	words []string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{nil}
}

func (this *MagicDictionary) BuildDict(dict []string) {
	this.words = append([]string{}, dict...)
}

func (this *MagicDictionary) Search(searchWord string) bool {
	cnt := make([]int8, len(this.words))

	for i := 0; i < len(searchWord); i++ {
		for j := 0; j < len(this.words); j++ {
			if len(this.words[j]) > len(searchWord) || len(this.words[j]) < len(searchWord) || cnt[j] > 1 {
				continue
			}
			if i < len(this.words[j]) {
				if this.words[j][i] != searchWord[i] {
					cnt[j]++
				}
				if i == len(searchWord)-1 && cnt[j] == 1 {
					return true
				}
			}
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

func main() {

	obj := Constructor()
	obj.BuildDict([]string{"hello", "leetcode"})
	println(obj.Search("leetcodd"))
}
