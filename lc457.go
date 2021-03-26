package main

import "fmt"

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		slow, fast := i, i
		direction := 1
		if nums[slow] < 0 {
			direction = -1
		}
		fmt.Printf("%+v\n", nums)
		println("slow:", slow, "fast:", fast)

		for {
			println("slow:", slow, nums[slow], "fast:", fast, nums[fast])

			slow = next(n, slow, nums[slow])
			if nums[slow]*direction <= 0 {
				break
			}

			fast = next(n, fast, nums[fast])
			if nums[fast]*direction <= 0 {
				break
			}

			fast = next(n, fast, nums[fast])

			if nums[fast]*direction <= 0 {
				break
			}
			if slow == fast {
				if slow == next(n, slow, nums[slow]) {
					break
				}
				return true
			}
		}
		slow = i
		for nums[slow]*direction > 0 {
			tmp := next(n, slow, nums[slow])
			nums[slow] = 0
			slow = tmp
		}
	}
	return false
}

func next(n, currPos, step int) int {
	return (n + currPos + step%n) % n
}

func main() {
	// nums := []int{2, -1, 1, 2, 2}
	// println(circularArrayLoop(nums))

	// nums = []int{-2, 1, -1, -2, -2}
	// println(circularArrayLoop(nums))

	// nums := []int{-1, 2}
	// println(circularArrayLoop(nums))

	// nums := []int{3, 1, 2}
	// nums := []int{1, -1, 1, 1}
	// nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, -5}
	// println(circularArrayLoop(nums))

	nums := []int{-2, -3, -9}
	println(circularArrayLoop(nums))
}
