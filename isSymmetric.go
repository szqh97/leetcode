package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	return r1.Val == r2.Val && isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root, root)
}
func createTree() *TreeNode {
	t1 := &TreeNode{Val: 1}
	tl2 := &TreeNode{Val: 2}
	tr2 := &TreeNode{Val: 2}
	tl3 := &TreeNode{Val: 3}
	tr3 := &TreeNode{Val: 3}
	tl4 := &TreeNode{Val: 4}
	tr4 := &TreeNode{Val: 4}
	t1.Left = tl2
	t1.Right = tr2
	tl2.Left = tl3
	tl2.Right = tl4
	tr2.Right = tr3
	tr2.Left = tr4
	return t1
}
func main() {
	fmt.Println(isSymmetric(createTree()))
	fmt.Println(isSymmetric(&TreeNode{Val: 1}))

}
