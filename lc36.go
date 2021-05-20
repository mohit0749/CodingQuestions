package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	mpRows := make(map[int]map[byte]map[string]bool)
	mpCols := make(map[int]map[byte]map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if string(board[i][j]) == `.` {
				continue
			}
			if mpRows[i] == nil {
				mpRows[i] = make(map[byte]map[string]bool)
			}
			if mpCols[j] == nil {
				mpCols[j] = make(map[byte]map[string]bool)
			}
			if mpRows[i][board[i][j]] == nil {
				mpRows[i][board[i][j]] = make(map[string]bool)
			}
			if mpCols[j][board[i][j]] == nil {
				mpCols[j][board[i][j]] = make(map[string]bool)
			}
			mpRows[i][board[i][j]][fmt.Sprintf("%d_%d", i, j)] = true
			mpCols[j][board[i][j]][fmt.Sprintf("%d_%d", i, j)] = true
		}
	}
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[i]); j += 3 {
			gridMp := make(map[byte]bool)
			for ii := i; ii < i+3; ii++ {
				for jj := j; jj < len(board[i]) && jj < j+3; jj++ {
					if string(board[ii][jj]) == "." {
						continue
					}
					if v, ok := mpRows[ii][board[ii][jj]]; ok && len(v) >= 2 {
						return false
					}
					if v, ok := mpCols[jj][board[ii][jj]]; ok && len(v) >= 2 {
						return false
					}
					if gridMp[board[ii][jj]] {
						return false
					}
					gridMp[board[ii][jj]] = true
				}
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	println(isValidSudoku(board))
}
