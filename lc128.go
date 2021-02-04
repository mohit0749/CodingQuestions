package main

func getMax(mp map[int]int) int {
	mx := 0
	for _, v := range mp {
		if mx < v {
			mx = v
		}
	}
	return mx
}

func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]-1]; !ok {
			var tmpNum = nums[i]
			seqNo := 1
			_, ok := mp[tmpNum+1]
			for ; ok; _, ok = mp[tmpNum+1] {
				tmpNum++
				seqNo++
			}
			if maxLen < seqNo {
				maxLen = seqNo
			}
		}
	}
	return maxLen
}

func main() {
	nums := []int{0,3,7,2,5,8,4,6,0,1}
	println(longestConsecutive(nums))
}
