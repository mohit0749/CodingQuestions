package main

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	var ans int
	for i := 0; i < len(arr); i++ {
		ans ^= arr[i]
	}
	println(ans)
	println("------")
	for i := 0; i < len(arr); i++ {
		ans ^= arr[i]
		if ans == 0 {
			println("--", arr[i])
		}
		println(ans)
	}
}
