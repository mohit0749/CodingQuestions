package main

type MapSum struct {
	exists map[string]int
	trie TrieNode
}


/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{make(map[string]int),TrieNode{make(map[byte]TrieNode),0}}
}


func (this *MapSum) Insert(key string, val int)  {
	oldVal,_:=this.exists[key]
	childNode:=this.trie
	for i:=0;i<len(key);i++{
		nextNode,ok:=childNode.childrens[key[i]]
		if !ok{
			nextNode.childrens=make(map[byte]TrieNode)
		}
		nextNode.val+=val-oldVal
		childNode.childrens[key[i]]=nextNode
		childNode=childNode.childrens[key[i]]
	}
	this.exists[key]=val
}


func (this *MapSum) Sum(prefix string) int {
	childNode:=this.trie
	for i:=0;i<len(prefix);i++{
		childNode=childNode.childrens[prefix[i]]
	}
	return childNode.val
}


type TrieNode struct{
	childrens map[byte]TrieNode
	val int
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
func main(){
	obj := Constructor()
	obj.Insert("apple",3)
	param_2 := obj.Sum("ap")
	println(param_2)
	obj.Insert("ap",2)
	println(obj.Sum("ap"))
}
