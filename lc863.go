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

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	graph := make(map[int][]int)
	buildGraph(nil, root, graph)
	bfs := make([]int, 1)
	bfs[0] = target.Val
	visited := make(map[int]bool)
	for i := 0; i < K; i++ {
		newLevel := make([]int, 0)
		for _, v := range bfs {
			for _, adj := range graph[v] {
				if _, ok := visited[adj]; !ok {
					newLevel = append(newLevel, adj)
				}
			}
			visited[v] = true
		}
		bfs = newLevel
	}
	return bfs
}

func buildGraph(root, child *TreeNode, mp map[int][]int) {
	if root != nil && child != nil {
		if mp[root.Val] == nil {
			mp[root.Val] = make([]int, 0)
			mp[root.Val] = append(mp[root.Val], child.Val)
		} else {
			mp[root.Val] = append(mp[root.Val], child.Val)
		}
		if mp[child.Val] == nil {
			mp[child.Val] = make([]int, 0)
			mp[child.Val] = append(mp[child.Val], root.Val)
		} else {
			mp[child.Val] = append(mp[child.Val], root.Val)
		}
	}
	if child == nil {
		return
	}
	if child.Left != nil {
		buildGraph(child, child.Left, mp)
	}
	if child.Right != nil {
		buildGraph(child, child.Right, mp)
	}
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 22,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{7, nil, nil},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{21, nil, nil},
			},
			Right: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   12,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   11,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: nil,
	}
	distanceK(root, root.Left, 2)
}
