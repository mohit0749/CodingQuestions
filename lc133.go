package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNode := &Node{node.Val, nil}
	visited := make(map[int]*Node)
	dfs(node, newNode, visited)
	return newNode
}

func dfs(node, newNode *Node, visited map[int]*Node) {
	visited[node.Val] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		visitedNode, ok := visited[node.Neighbors[i].Val]
		if !ok {
			tmpNode := &Node{node.Neighbors[i].Val, nil}
			newNode.Neighbors = append(newNode.Neighbors, tmpNode)
			dfs(node.Neighbors[i], tmpNode, visited)
		} else {
			newNode.Neighbors = append(newNode.Neighbors, visitedNode)
		}
	}
}
