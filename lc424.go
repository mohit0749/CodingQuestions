package main

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	windowMap := make(map[uint8]int)
	l := 0
	windowMap[s[l]]++
	max := 1
	requiredOp:=0
	for r := l + 1; r < len(s); {
		windowMap[s[r]]++
		if requiredOp<windowMap[s[r]]{
			requiredOp=windowMap[s[r]]
		}
		// requiredOp = getCount(windowMap)
		for r-l+1-requiredOp > k {
			windowMap[s[l]]--
			l++
		}
		if max < r-l+1 {
			max = r - l + 1
		}
		r++
	}
	return max
}

func getCount(mp map[uint8]int) int {
	maxCnt := 0
	var maxKey uint8
	for k, v := range mp {
		if v > maxCnt {
			maxCnt = v
			maxKey = k
		}
	}
	op := 0
	for k, v := range mp {
		if k != maxKey {
			op += v
		}
	}
	return op
}

func main() {
	s := "ABAB"
	k := 1
	//s:="AABCBBA"
	//k:=2
	println(characterReplacement(s, k))
}
