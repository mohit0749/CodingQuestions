package main

import "codes/src/treebuidler"

func printBottomView(root *treebuidler.TreeNode) {
	mp := make(map[int]int)
	bottomView(root, 0, mp)
	for _, v := range mp {
		println(v)
	}
}
func bottomView(root *treebuidler.TreeNode, dist int, verticalDist map[int]int) {
	if root == nil {
		return
	}
	bottomView(root.Left, dist-1, verticalDist)
	if _, ok := verticalDist[dist]; !ok {
		verticalDist[dist] = root.Val
	}
	bottomView(root.Right, dist+1, verticalDist)

}

func main() {
	root := treebuidler.BuildTree([]interface{}{20, 8, 22, 5, 3, nil, 25, nil, nil, 10, 14})
	//root := treebuidler.BuildTree([]interface{}{1, 3, 2})
	printBottomView(root)
}
