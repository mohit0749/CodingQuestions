package main

func binarySearch(arr []int, value, l, r int) int {
	for l <= r {
		var mid int = (l + r) / 2
		if arr[mid] == value {
			return mid
		}
		if value < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func UpperBound(arr []int, value, l, r int) int {
	for l <= r {
		var mid int = (l + r) / 2
		if arr[mid] == value {
			return mid
		}
		if value < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l > len(arr)-1 {
		return len(arr) - 1
	}
	return l
}

func searchMatrix(matrix [][]int, target int) bool {
	cols := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		cols = append(cols, matrix[i][len(matrix[i])-1])
	}
	row := UpperBound(cols, target, 0, len(cols)-1)
	if binarySearch(matrix[row], target, 0, len(matrix[row])-1) != -1 {
		return true
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	println(searchMatrix(matrix, 2))
	println(searchMatrix(matrix, 70))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			println(searchMatrix(matrix, matrix[i][j]))
		}
	}

}
