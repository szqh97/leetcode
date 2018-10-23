package main

import "fmt"

/*
*  给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
*  你可以假设除了数字 0 之外，这两个数字都不会以零开头。
*  示例：
*  输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
*  输出：7 -> 0 -> 8
*  原因：342 + 465 = 807
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() {
	for l != nil {
		fmt.Printf("%d", l.Val)
		l = l.Next
		if l != nil {
			fmt.Printf("->")
		}
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 4}
	//l3 := &ListNode{Val: 3}
	l1.Next = l2
	//l2.Next = l3

	r1 := &ListNode{Val: 5}
	r2 := &ListNode{Val: 6}
	r3 := &ListNode{Val: 4}
	r1.Next = r2
	r2.Next = r3
	l1.String()
	r1.String()

	result := addTwoNumbers(l1, r1)

	result.String()

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	it1 := l1
	it2 := l2
	t := 0
	var h, pos *ListNode = nil, nil
	for it1 != nil || it2 != nil {
		var x int = 0
		if it1 != nil {
			x = it1.Val
		}

		var y int = 0
		if it2 != nil {
			y = it2.Val
		}

		sum := x + y + t
		t = sum / 10
		n := &ListNode{Val: sum % 10, Next: nil}
		if h == nil {
			h = n
			pos = h
		} else {
			pos.Next = n
			pos = pos.Next
		}
		if it1 != nil {
			it1 = it1.Next
		}
		if it2 != nil {
			it2 = it2.Next
		}

	}
	if t != 0 {
		pos.Next = &ListNode{Val: 1}
	}
	return h
}
