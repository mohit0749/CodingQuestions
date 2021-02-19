package main

import "fmt"

func findWords(board [][]byte, words []string) []string {

	ans := make([]string, 0)
	for _, word := range words {
		visited := make([][]bool, 0)
		for i := 0; i < len(board); i++ {
			visited = append(visited, make([]bool, len(board[i])))
		}
		coords := traverseMat(board, word[0])
		isFound := false
		for _, coord := range coords {
			if dfs(coord[0], coord[1], board, visited, word, 1) {
				isFound = true
				break
			}
		}
		if isFound {
			ans = append(ans, word)
		}
	}
	return ans
}

func traverseMat(board [][]byte, c byte) [][2]int {
	coords := make([][2]int, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if c == board[i][j] {
				coords = append(coords, [2]int{i, j})
			}
		}
	}
	return coords
}

func dfs(r, c int, board [][]byte, visited [][]bool, word string, i int) bool {
	visited[r][c] = true
	if i >= len(word) {
		return true
	}
	found := false
	if r+1 < len(board) && !visited[r+1][c] && word[i] == board[r+1][c] {
		found = found || dfs(r+1, c, board, visited, word, i+1)
	}
	if r-1 >= 0 && !visited[r-1][c] && word[i] == board[r-1][c] {
		found = found || dfs(r-1, c, board, visited, word, i+1)
	}

	if c+1 < len(board[r]) && !visited[r][c+1] && word[i] == board[r][c+1] {
		found = found || dfs(r, c+1, board, visited, word, i+1)
	}

	if c-1 >= 0 && !visited[r][c-1] && word[i] == board[r][c-1] {
		found = found || dfs(r, c-1, board, visited, word, i+1)

	}
	visited[r][c] = false
	return found
}

func main() {
	//board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	//words := []string{"oath", "pea", "eat", "rain"}

	board := [][]byte{{'a', 'b'}, {'c', 'd'}}
	words := []string{"abcb"}
	//board := [][]byte{
	//	{'a', 'b', 'c'},
	//	{'a', 'e', 'd'},
	//	{'a', 'f', 'g'}}
	//words := []string{"abcdefg", "befa", "eaabcdgfa", "gfedcbaaa"}
	//["abcdefg","befa","eaabcdgfa","gfedcbaaa"]
	fmt.Printf("%+v\n", findWords(board, words))
}
