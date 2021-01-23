package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if recurFindWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func recurFindWord(board [][]byte, word string, charAt, i, j int) bool {
	if charAt >= len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || word[charAt] != board[i][j] {
		return false
	}

	temp := board[i][j]
	board[i][j] = '*'
	ret := recurFindWord(board, word, charAt+1, i+1, j) ||
		recurFindWord(board, word, charAt+1, i, j+1) ||
		recurFindWord(board, word, charAt+1, i-1, j) ||
		recurFindWord(board, word, charAt+1, i, j-1)
	board[i][j] = temp
	return ret
}

func main() {
	w := "ABCCED"
	board := [][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}

	println(exist(board, w))
}
