package main

import "fmt"

func main() {
	//island := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	island := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Print(numIslands((island)))
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j, grid)
				count += 1
			}
		}
	}
	return count
}
func dfs(i, j int, grid [][]byte) {
	grid[i][j] = 0
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		dfs(i+1, j, grid)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
		dfs(i, j+1, grid)
	}
	if i-1 >= 0 && grid[i-1][j] == 1 {
		dfs(i-1, j, grid)
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		dfs(i, j-1, grid)
	}
}
