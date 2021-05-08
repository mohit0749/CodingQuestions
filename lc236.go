package main

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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode
	LCA(root, p, q, &lca)
	return lca
}

func LCA(root, p, q *TreeNode, lca **TreeNode) bool {
	if root == nil {
		return false
	}
	if *lca != nil {
		return true
	}
	lf := LCA(root.Left, p, q, lca)
	rf := LCA(root.Right, p, q, lca)
	if lf && rf {
		if *lca == nil {
			*lca = root
		}
		return true
	}
	var found bool
	if root.Val == p.Val || root.Val == q.Val {
		found = true
	}
	if (found && lf) || (found && rf) {
		if *lca == nil {
			*lca = root
		}
		return true
	}
	if lf {
		return true
	}
	if rf {
		return true
	}
	if found {
		return true
	}

	return false
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	lca := lowestCommonAncestor(root, &TreeNode{5, nil, nil}, &TreeNode{4, nil, nil})
	println(lca.Val)
}

/*
[3,5,1,6,2,0,8,null,null,7,4]
5
1
[3,5,1,6,2,0,8,null,null,7,4]
6
7
[3,5,1,6,2,0,8,null,null,7,4]
6
4
[3,5,1,6,2,0,8,null,null,7,4]
7
4
[3,5,1,6,2,0,8,null,null,7,4]
5
2
[3,5,1,6,2,0,8,null,null,7,4]
6
1
[3,5,1,6,2,0,8,null,null,7,4]
0
8
[3,5,1,6,2,0,8,null,null,7,4]
7
1
*/
