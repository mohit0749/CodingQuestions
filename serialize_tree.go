package main

import (
	"codes/src/treebuidler"
	"fmt"
)

func serialize(root *treebuidler.TreeNode, arr *[]interface{}) {
	if root == nil {
		*arr = append(*arr, nil)
		return
	}
	*arr = append(*arr, root.Val)
	serialize(root.Left, arr)
	serialize(root.Right, arr)
}

func SerializeTree(root *treebuidler.TreeNode) []interface{} {
	arr := make([]interface{}, 0)
	serialize(root, &arr)
	return arr
}

func main() {
	root := treebuidler.BuildTree([]interface{}{20, 8, nil, 4, 12, nil, nil, nil, nil, 10, 14, nil, nil, nil, nil})
	fmt.Printf("%+v\b", SerializeTree(root))
}
