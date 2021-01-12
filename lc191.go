package main

func hammingWeight(num uint32) int {
	cnt := 0
	for num > 0 {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}
	return cnt
}

func main() {
	var n uint32 = 0b00000000000000000000000000001011
	println(n)
	println(hammingWeight(n))
}
