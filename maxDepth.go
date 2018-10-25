package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func createTree() *TreeNode {
	t1 := &TreeNode{Val: 3}
	t9 := &TreeNode{Val: 9}
	t20 := &TreeNode{Val: 20}
	t15 := &TreeNode{Val: 15}
	t7 := &TreeNode{Val: 7}
	t1.Left = t9
	t1.Right = t20
	t20.Left = t15
	t20.Right = t7
	return t1

}
func main() {
	fmt.Println(maxDepth(createTree()))

}
