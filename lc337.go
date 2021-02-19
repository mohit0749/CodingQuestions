package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	robbed := maximizeRob(root)
	return max(robbed[0], robbed[1])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maximizeRob(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}
	left := maximizeRob(root.Left)
	right := maximizeRob(root.Right)
	rob := root.Val + left[1] + right[1]
	notRob := max(left[0], left[1]) + max(right[0], right[1])
	return [2]int{rob, notRob}
}

func main() {

}
