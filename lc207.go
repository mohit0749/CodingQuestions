package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	graph := make(map[int][]int)
	for _, v := range prerequisites {
		graph[v[0]] = append(graph[v[0]], v[1])
	}

	visited := make([]bool, numCourses)
	recur := make([]bool, numCourses)
	for k, _ := range graph {
		if !visited[k] {
			if dfs(k, graph, visited, recur){
				return false
			}
		}
	}
	return true
}

func dfs(src int, graph map[int][]int, visited,recur []bool) bool {
	visited[src] = true
	recur[src]=true
	for _, v := range graph[src] {
		if !visited[v] {
			if dfs(v, graph, visited, recur){
				return true
			}
		} else if recur[v]==true {
			return true
		}
	}
	recur[src]=false
	return false
}

func main() {
	numCourses := 20
	prerequisites := [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}
	//numCourses:=5
	//prerequisites:=[][]int{{1,4},{2,4},{3,1},{3,2}}
	//prerequisites := [][]int{{0, 1}, {1, 0}}
	println(canFinish(numCourses, prerequisites))
}
