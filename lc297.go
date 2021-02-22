package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	return fmt.Sprintf("%d,%s,%s", root.Val, left, right)
}

//func _serialize(root *TreeNode, s string) string {
//
//}
var _deserialize func(root *TreeNode, nodes []*TreeNode) *TreeNode

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := make([]*TreeNode, 0)
	for _, s := range strings.Split(data, ",") {
		if s == "null" {
			arr = append(arr, nil)
		} else {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, &TreeNode{n, nil, nil})
		}
	}
	if len(arr) == 0 {
		return nil
	}
	root := arr[0]
	i := 0
	_deserialize = func(root *TreeNode, nodes []*TreeNode) *TreeNode {
		if i >= len(nodes) {
			return nil
		}
		i = i + 1
		if root == nil {
			return nil
		}

		root.Left = _deserialize(nodes[i], nodes)
		root.Right = _deserialize(nodes[i], nodes)
		return root
	}
	return _deserialize(root, arr)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {
	codec := Constructor()
	root := &TreeNode{3, &TreeNode{4, &TreeNode{6, nil, nil}, nil},
		&TreeNode{5, &TreeNode{7, nil, nil}, nil}}
	data := codec.serialize(root)
	println(data)
	newRoot := codec.deserialize(data)
	println(codec.serialize(newRoot))
}
