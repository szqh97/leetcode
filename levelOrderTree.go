package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	NodeSlice := []*TreeNode{}
	result := [][]int{}
	if root == nil {
		return result
	}
	NodeSlice = append(NodeSlice, root)
	NodeSlice = append(NodeSlice, nil)
	level := []int{}
	for i := 0; i < len(NodeSlice); i++ {
		if NodeSlice[i] != nil {
			level = append(level, NodeSlice[i].Val)
			if NodeSlice[i].Left != nil {
				NodeSlice = append(NodeSlice, NodeSlice[i].Left)
			}
			if NodeSlice[i].Right != nil {
				NodeSlice = append(NodeSlice, NodeSlice[i].Right)
			}
		}
		if NodeSlice[i] == nil {
			if i >= 1 && NodeSlice[i-1] != nil {
				NodeSlice = append(NodeSlice, nil)
				result = append(result, level)
				level = []int{}
			}
		}
	}
	return result
}
func main() {
	root := &TreeNode{Val: 9}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(root))

}
