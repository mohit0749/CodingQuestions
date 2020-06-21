package main

import "fmt"

func main() {
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(plusOne(d))
}

func plusOne(digits []int) []int {
	var carry int = 1

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
		if carry == 0 {
			break
		}
	}
	if carry > 0 {
		temp := []int{1}
		temp = append(temp, digits...)
		return temp
	}
	return digits
}
