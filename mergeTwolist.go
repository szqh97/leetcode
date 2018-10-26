package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {

	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l *ListNode = nil
	var l0 *ListNode = nil
	for l1 != nil && l2 != nil {

		n := l1
		if l1.Val >= l2.Val {
			n = l2
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
		if l == nil {
			l = n
			l0 = n
		} else {
			l.Next = n
			l = l.Next
		}

	}

	if l1 == nil && l2 != nil {
		l.Next = l2
	}
	if l2 == nil && l1 != nil {
		l.Next = l1
	}
	return l0
}
func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l4

	r1 := &ListNode{Val: 1}
	r3 := &ListNode{Val: 3}
	r4 := &ListNode{Val: 4}
	r1.Next = r3
	r3.Next = r4

	m := mergeTwoLists(l1, r1)
	m.Print()
}
