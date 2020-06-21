package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	i := 0
	for i < len(flowerbed) {
		if flowerbed[i] == 0 {
			if (i == 0 && len(flowerbed) == 1) || (i == 0 && i+1 < len(flowerbed) && flowerbed[i+1] == 0 || (i == len(flowerbed)-1 && flowerbed[i-1] == 0)) ||
				(i-1 > 0 && flowerbed[i-1] == 0 && i+1 < len(flowerbed) && flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				n--
				if n == 0 {
					break
				}
				i += 2
				continue
			}
		}
		i++
	}
	if n == 0 {
		return true
	}
	return false
}

func main() {
	ar := []int{0}
	n := 1
	fmt.Println(canPlaceFlowers(ar, n))
}
