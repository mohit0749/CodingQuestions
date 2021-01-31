package main

import "codes/src/treebuidler"

func convertToDLL(root *treebuidler.TreeNode) *treebuidler.TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	var leftHead, rightHead *treebuidler.TreeNode
	if root.Left != nil {
		leftHead = convertToDLL(root.Left)
	}
	if root.Right != nil {
		rightHead = convertToDLL(root.Right)
	}
	root.Left = getTail(leftHead)
	if root.Left != nil {
		root.Left.Right = root
	}
	root.Right = getHead(rightHead)
	if root.Right != nil {
		root.Right.Left = root
	}
	return root
}

func getTail(root *treebuidler.TreeNode) *treebuidler.TreeNode {
	if root == nil {
		return root
	}
	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}
	return temp
}

func getHead(root *treebuidler.TreeNode) *treebuidler.TreeNode {
	if root == nil {
		return root
	}
	temp := root
	for temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func convertToDLLInorder(root *treebuidler.TreeNode, head, tail **treebuidler.TreeNode) {
	if root == nil {
		return
	}
	convertToDLLInorder(root.Left, head, tail)
	if *head == nil {
		*head = root
	}
	root.Left = *tail

	if *tail != nil {
		(*tail).Right = root
	}
	*tail = root

	convertToDLLInorder(root.Right, head, tail)
}

func main() {
	root := treebuidler.BuildTree([]interface{}{10, 12, 15, 25, 30, 36, nil})
	//convertToDLL(root)
	var head, tail *treebuidler.TreeNode
	convertToDLLInorder(root, &head, &tail)
	for head != nil {
		println(head.Val)
		head = head.Right
	}
}
