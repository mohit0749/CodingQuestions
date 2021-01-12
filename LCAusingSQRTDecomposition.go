package main

import "math"

/*
https://www.geeksforgeeks.org/sqrt-square-root-decomposition-set-2-lca-tree-osqrth-time/

The key concept here is that first we bring both the nodes in same group and
having same jumpParent by climbing decomposed blocks(block size of sqrt(h))
above the tree one by one and then when both the nodes are in same group
and have same jumpParent we use our naive approach to find LCA of the nodes.
*/
const MAXN = 1001

var blockSize int
var depth, parent, jumpParent []int
var adj map[int][]int

func Dfs(currNode, parentNode int) {
	depth[currNode] = depth[parentNode] + 1
	parent[currNode] = parentNode

	if depth[currNode]%blockSize == 0 {
		jumpParent[currNode] = parentNode
	} else {
		jumpParent[currNode] = jumpParent[parentNode]
	}

	for _, childNode := range adj[currNode] {
		Dfs(childNode, currNode)
	}
}

func preProcess(height int) {
	blockSize = int(math.Sqrt(float64(height)))
	depth = make([]int, MAXN)
	parent = make([]int, MAXN)
	jumpParent = make([]int, MAXN)

	depth[0] = -1
	Dfs(1, 0)
}

//it works only if both node are on the same depth
func LCANaive(u, v int) int {
	for u != v {
		if depth[v] < depth[u] {
			u, v = v, u
		}
		v = parent[v]
	}
	return v
}

//works by jumping in blocks
func LCASqrt(u, v int) int {
	for jumpParent[u] != jumpParent[v] {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		u = jumpParent[u]
	}
	return LCANaive(u, v)
}

func main() {
	adj = make(map[int][]int)
	// tree
	adj[1] = []int{2, 3, 4}
	adj[2] = []int{5}
	adj[3] = []int{6, 7}
	adj[4] = []int{8}
	adj[6] = []int{9, 10, 11}
	adj[8] = []int{12}
	adj[9] = []int{13, 14}
	adj[11] = []int{15}
	adj[12] = []int{16}
	adj[14] = []int{17, 18, 19}
	adj[15] = []int{20, 21}
	adj[17] = []int{22}
	adj[19] = []int{23, 24}
	adj[21] = []int{25, 26}
	adj[23] = []int{27, 28}
	adj[25] = []int{29, 30}
	adj[28] = []int{31, 32, 33}
	adj[30] = []int{34, 35}
	preProcess(9)
	println(LCASqrt(29, 9))
}
