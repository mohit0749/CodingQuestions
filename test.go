package main

func find_affected(arr []int, k int) int {
	l := 0
	r := 0
	sum := 0
	cnt := 0
	flag := false
	for r < len(arr) {
		if sum+arr[r] < k {
			if flag {
				cnt -= countRange(r - l)
			}
			flag = false
			sum += arr[r]
			r++
			if r == len(arr) {
				if sum < k {
					cnt += countRange(r - l)
				}
			}
			//cnt = false
			continue
		} else {
			if sum < k {
				cnt += countRange(r - l)
			}
			for l < r && sum+arr[r] > k {
				sum -= arr[l]
				l++
			}
			if l == r {
				sum += arr[r]
				r++
				if r == len(arr) {
					if sum < k {
						cnt += countRange(r - l)
					}
				}
			} else {
				flag = true
			}
		}
	}
	return cnt
}

func countRange(n int) int {
	if n <= 0 {
		return n
	}
	return (n * (n + 1)) / 2
}

func main() {
	//arr := []int{1, 3, 4, 10} //k=10
	//arr := []int{1, 11, 2, 3, 15} //,k=10
	arr := []int{2, 5, 6}
	k := 10
	println(find_affected(arr, k))
}
