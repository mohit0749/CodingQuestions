package main

import "codes/src/treebuidler"

func LCA(root *treebuidler.TreeNode, n1, n2 int) *treebuidler.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == n1 || root.Val == n2 {
		return root
	}
	left := LCA(root.Left, n1, n2)
	right := LCA(root.Right, n1, n2)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func LCABST(root *treebuidler.TreeNode, n1, n2 int) *treebuidler.TreeNode {
	if root == nil {
		return nil
	}
	if n1 < root.Val && n2 < root.Val {
		return LCABST(root.Left, n1, n2)
	} else if n1 > root.Val && n2 > root.Val {
		return LCABST(root.Right, n1, n2)
	} else if (n1 < root.Val && n2 > root.Val) || (n1 > root.Val && n2 < root.Val) {
		left := LCABST(root.Left, n1, n2)
		right := LCABST(root.Right, n1, n2)
		if left != nil && right != nil {
			return root
		}
	}
	return root
}

func main() {
	root := treebuidler.BuildTree([]interface{}{5, 4, 6, 3, nil, nil, 7, nil, nil, nil, nil, nil, nil, nil, 8})
	//LCA := LCA(root, 3, 6)
	LCA := LCABST(root, 4, 3)
	if LCA != nil {
		println(LCA.Val)
	}
}
