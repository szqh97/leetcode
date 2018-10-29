package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	Print(roto, &result)
	return result
}
func Print(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	Print(root.Left, result)
	Print(root.Right, result)
}
func main() {
	t := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t.Right = t2
	result := []int{}
	Print(t, &result)
	fmt.Println(result)

}
