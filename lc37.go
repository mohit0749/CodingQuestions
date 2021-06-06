package main

import "fmt"

var rowMap, colMap map[int]map[byte]bool
var gridMap map[string]map[byte]bool

func solveSudoku(board [][]byte) {
	rowMap = make(map[int]map[byte]bool)
	colMap = make(map[int]map[byte]bool)
	gridMap = make(map[string]map[byte]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if string(board[i][j]) == `.` {
				continue
			}
			if rowMap[i] == nil {
				rowMap[i] = make(map[byte]bool)
			}
			if colMap[j] == nil {
				colMap[j] = make(map[byte]bool)
			}
			gKey := fmt.Sprintf("gi%dgj%d", i/3, j/3)
			if gridMap[gKey] == nil {
				gridMap[gKey] = make(map[byte]bool)
			}
			rowMap[i][board[i][j]] = true
			colMap[j][board[i][j]] = true
			gridMap[gKey][board[i][j]] = true
		}
	}
	ii, jj := 0, 0
	for i := 0; i < len(board); i++ {
		isBreak := false
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				ii, jj = i, j
				isBreak = true
				break
			}
		}
		if isBreak {
			break
		}
	}
	solve(board, ii, jj)
}

func solve(board [][]byte, i, j int) bool {
	if i >= len(board) || j >= len(board[0]) {
		return true
	}
	for n := 1; n <= 9; n++ {
		currChar := '0' + byte(n)
		if rowMap[i] == nil {
			rowMap[i] = make(map[byte]bool)
		}
		if colMap[j] == nil {
			colMap[j] = make(map[byte]bool)
		}
		gKey := fmt.Sprintf("gi%dgj%d", i/3, j/3)
		if gridMap[gKey] == nil {
			gridMap[gKey] = make(map[byte]bool)
		}
		if rowMap[i][currChar] {
			continue
		}
		if colMap[j][currChar] {
			continue
		}
		if gridMap[gKey][currChar] {
			continue
		}
		rowMap[i][currChar] = true
		colMap[j][currChar] = true
		gridMap[gKey][currChar] = true
		board[i][j] = currChar
		isBreak := false
		for ii := 0; ii < len(board); ii++ {
			for jj := 0; jj < len(board); jj++ {
				if board[ii][jj] == '.' {
					if !solve(board, ii, jj) {
						isBreak = true
						break
					}
				}
			}
			if isBreak {
				break
			}
		}
		if !isBreak {
			return true
		}
		rowMap[i][board[i][j]] = false
		colMap[j][board[i][j]] = false
		gridMap[gKey][board[i][j]] = false
		board[i][j] = '.'
	}
	return false
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			print(string(board[i][j]))
			print(",")
		}
		println()
	}
}
