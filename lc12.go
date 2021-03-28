package main

import (
	"strconv"
	"strings"
)

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.



Example 1:

Input: num = 3
Output: "III"
Example 2:

Input: num = 4
Output: "IV"
Example 3:

Input: num = 9
Output: "IX"
Example 4:

Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 5:

Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:

1 <= num <= 3999

*/

func intToRoman(num int) string {
	mp := getRomanMapping()
	digits := generateNumber(num)
	var ans strings.Builder
	for i := len(digits) - 1; i >= 0; i-- {
		ans.WriteString(getRomanNumber(digits[i], mp))
	}
	return ans.String()
}

func getRomanMapping() map[int]string {
	mp := make(map[int]string)
	mp[1] = "I"
	mp[4] = "IV"
	mp[5] = "V"
	mp[9] = "IX"
	mp[10] = "X"
	mp[40] = "XL"
	mp[50] = "L"
	mp[90] = "XC"
	mp[100] = "C"
	mp[400] = "CD"
	mp[500] = "D"
	mp[900] = "CM"
	mp[1000] = "M"
	return mp
}

func getRomanNumber(n int, mp map[int]string) string {
	if v, ok := mp[n]; ok {
		return v
	}
	var s strings.Builder
	sub := len(strconv.Itoa(n))
	pos := 1
	for i := 1; i < sub; i++ {
		pos *= 10
	}
	//println(sub)
	for n > 0 {
		if v, ok := mp[n]; ok {
			return v + s.String()
		}
		s.WriteString(mp[pos])
		n -= pos
	}
	return s.String()
}

func generateNumber(n int) []int {
	nums := make([]int, 0)
	cnt := 1
	for n > 0 {
		nums = append(nums, (n%10)*cnt)
		cnt *= 10
		n /= 10
	}
	return nums
}

func main() {
	//n := 1234
	//println(intToRoman(n))
	for i := 1; i < 4000; i++ {
		println(i)
	}
}
