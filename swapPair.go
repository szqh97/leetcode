package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}
func swapPairs(head *ListNode) *ListNode {
	var h *ListNode = nil
	var p, t *ListNode
	p = head
	if p == nil || p.Next == nil {
		return p
	}
	for p != nil && p.Next != nil {
		q := p.Next
		if q == nil && h == nil {
			return p
		}
		p.Next = q.Next
		q.Next = p
		if h == nil {
			h = q
		}
		if t == nil {
			t = p
		} else {
			t.Next = q
			t = p
		}
		p = p.Next

	}
	return h
}

func main() {
	l1 := &ListNode{Val: 1}
	//l1.Next = &ListNode{Val: 2}
	//l1.Next.Next = &ListNode{Val: 3}
	//l1.Next.Next.Next = &ListNode{Val: 4}

	r := swapPairs(l1)
	r.print()

}
