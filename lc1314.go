package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("%+v", matrixBlockSum(mat, 1))
}

func cumulativeSUm(mat [][]int) [][]int {
	var cum [][]int
	cum = make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		cum[i] = append(cum[i], make([]int, len(mat[i]))...)
		for j := 0; j < len(mat[i]); j++ {
			sum := mat[i][j]
			if i-1 >= 0 {
				sum += cum[i-1][j]
			}
			if j-1 >= 0 {
				sum += cum[i][j-1]
			}
			if i-1 >= 0 && j-1 >= 0 {
				sum -= cum[i-1][j-1]
			}
			cum[i][j] = sum
		}
	}
	return cum
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	cnum := cumulativeSUm(mat)
	bmat := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		bmat[i] = append(bmat[i], make([]int, len(mat[i]))...)
		for j := 0; j < len(mat[i]); j++ {
			sum := 0
			tmp_i, tmp_j := i, j
			tmp_i += k
			if tmp_i > len(mat)-1 {
				tmp_i = len(mat) - 1
			}
			tmp_j += k
			if tmp_j > len(mat[i])-1 {
				tmp_j = len(mat[i]) - 1
			}
			sum += cnum[tmp_i][tmp_j]

			if i-1-k >= 0 && j-1-k >= 0 {
				sum += cnum[i-1-k][j-1-k]
			}
			if j-1-k >= 0 {
				sum -= cnum[tmp_i][j-1-k]
			}
			if i-1-k >= 0 {
				sum -= cnum[i-1-k][tmp_j]
			}
			bmat[i][j] = sum
		}
	}
	return bmat
}
