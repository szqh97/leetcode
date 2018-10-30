package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(height(root.Left), height(root.Right)) + 1
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	/*
		root.Left.Left = &TreeNode{Val: 3}
		root.Right.Right = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 4}
		root.Right.Right.Right = &TreeNode{Val: 4}

	*/
	fmt.Println(isBalanced(root))

}
