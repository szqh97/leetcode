package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Print() {
	if t == nil {
		return
	}
	t.Left.Print()
	fmt.Printf("%d ", t.Val)
	t.Right.Print()
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(inorder)
	if l == 0 && l == 0 {
		return nil
	}
	if l == 1 && l == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	i := 0
	for ; i < l; i++ {
		if inorder[i] == postorder[l-1] {
			break
		}
	}
	root := &TreeNode{Val: postorder[l-1]}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:l])
	return root

}
func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	/*
		inorder := []int{3, 2, 1}
		postorder := []int{3, 2, 1}
	*/
	t := buildTree(inorder, postorder)
	t.Print()

}
