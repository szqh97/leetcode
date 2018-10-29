package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	fmt.Println(i)

	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}
func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, inorder))

}
