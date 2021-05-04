package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	stack := make([]*TreeNode, 0)
	level := 0
	stack = append(stack, root)
	zigzag := make([][]int, 0)
	for len(stack) > 0 {

		arrNode := make([]*TreeNode, 0)
		for len(stack) > 0 {
			item := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if item == nil {
				if level == 0 {
					level = 1
				} else {
					level = 0
				}
				break
			}
			arrNode = append(arrNode, item)
		}
		order := make([]int, 0)
		for i := 0; i < len(arrNode); i++ {
			order = append(order, arrNode[i].Val)
			if level%2 == 0 {
				if arrNode[i].Left != nil {
					stack = append(stack, arrNode[i].Left)
				}
				if arrNode[i].Right != nil {
					stack = append(stack, arrNode[i].Right)
				}
			} else {
				if arrNode[i].Right != nil {
					stack = append(stack, arrNode[i].Right)
				}
				if arrNode[i].Left != nil {
					stack = append(stack, arrNode[i].Left)
				}
			}
		}
		if len(order) > 0 {
			stack = append(stack, nil)
			zigzag = append(zigzag, order)
		}
	}
	return zigzag
}

func main() {
	//[3,9,20,1,2,15,7,11,null,21,22]
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{1, &TreeNode{11, nil, nil}, &TreeNode{12, nil, nil}}, Right: &TreeNode{2, nil, nil}},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{15, &TreeNode{51, nil, nil}, nil},
			Right: &TreeNode{7, nil, &TreeNode{71, nil, nil}},
		}}
	zigzagLevelOrder(tree)
}
