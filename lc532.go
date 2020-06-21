package main

import "fmt"
import "math"

type pair struct {
	a, b int
}

func findPairs(nums []int, k int) int {
	mp := make(map[int]int, 0)
	for _, v := range nums {
		mp[v] += 1
	}
	c := 0
	uniqMap := make(map[pair]int)
	for _, v := range nums {
		tmp := v - k
		if count, ok := mp[tmp]; ok {
			if tmp == v && count <= 1 {
				continue
			}
			t := int(math.Abs(float64(v - tmp)))
			if t != k {
				continue
			}
			var p1, p2 pair
			p1.a = v
			p1.b = tmp
			p2.a = tmp
			p2.b = v
			_, ok1 := uniqMap[p1]
			_, ok2 := uniqMap[p2]
			if ok1 || ok2 {
				continue
			}
			if int(math.Abs(float64(v-tmp))) == k {
				c++
				uniqMap[p1] += 1
			}

		}
	}
	fmt.Println(uniqMap)
	return c
}

func main() {
	nums := []int{1, 3, 1, 5, 4}
	k := 1
	fmt.Println(findPairs(nums, k))
}
