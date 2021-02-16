package main

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	totalPalinCnt := 0

	//extend for odd length
	for i := 0; i < len(s); i++ {
		l := i - 1
		r := i + 1
		cnt := 1
		for l >= 0 && r < len(s) &&  s[l] == s[r] {
			cnt++
			r++
			l--
		}
		totalPalinCnt += cnt
	}

	//extend for even length
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			l := i - 2
			r := i + 1
			cnt := 1
			for l >= 0 && r < len(s) && s[l] == s[r]{
				cnt++
				r++
				l--
			}
			totalPalinCnt += cnt
		}
	}
	return totalPalinCnt
}

func main() {
	s := "aaabaaa"
	println(countSubstrings(s))
}
