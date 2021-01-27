package main

func equilibPoint(arr []int) int {
	rightSum := 0
	for i := len(arr) - 1; i >= 0; i-- {
		rightSum += arr[i]
	}
	leftSum := 0
	for i := 0; i < len(arr); i++ {
		leftSum += arr[i]
		rem := rightSum - (leftSum - arr[i])
		if leftSum == rem {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 2, 2}
	println(equilibPoint(arr))
}
