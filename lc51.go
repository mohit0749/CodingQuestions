package main

import "fmt"

func solveNQueens(n int) [][]string {
	results := make([][]string, 0)
	board := make([][]byte, 0)
	for i := 0; i < n; i++ {
		row := make([]byte, 0)
		for j := 0; j < n; j++ {
			row = append(row, '.')
		}
		board = append(board, row)
	}
	backtrack(n, 0, n, board, &results)
	return results
}

func backtrack(n int, startrow, qcnt int, board [][]byte, results *[][]string) {
	if qcnt == 0 {
		s := make([]string, 0)
		for i := 0; i < len(board); i++ {
			s = append(s, string(board[i]))
		}
		fmt.Println(s)
		*results = append(*results, s)
		return
	}
	for i := startrow; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' && checkDiagonal(n, i, j, board) && checkColumn(n, i, j, board) && checkRow(n, i, j, board) {
				board[i][j] = 'Q'
				backtrack(n, i+1, qcnt-1, board, results)
				board[i][j] = '.'
			}
		}
	}
}

func checkDiagonal(n, i, j int, board [][]byte) bool {
	for ii, jj := i-1, j-1; ii >= 0 && jj >= 0; ii, jj = ii-1, jj-1 {
		if board[ii][jj] == 'Q' {
			return false
		}
	}
	for ii, jj := i-1, j+1; ii >= 0 && jj < n; ii, jj = ii-1, jj+1 {
		if board[ii][jj] == 'Q' {
			return false
		}
	}
	return true
}

func checkColumn(n, i, j int, board [][]byte) bool {
	for ii := i - 1; ii >= 0; ii-- {
		if board[ii][j] == 'Q' {
			return false
		}
	}
	return true
}

func checkRow(n, i, j int, board [][]byte) bool {
	for jj := j - 1; jj >= 0; jj-- {
		if board[i][jj] == 'Q' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solveNQueens(4))
}
