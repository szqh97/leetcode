package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for l != nil {
		fmt.Printf("%d", l.Val)
		l = l.Next
		if l != nil {
			fmt.Printf("->")
		}
	}
	fmt.Println()

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	cnt := 0
	for q.Next != nil {
		cnt = cnt + 1
		q = q.Next
		if cnt == n {
			break
		}
	}

	if cnt == n-1 {
		l := head
		head = head.Next
		l.Next = nil
		return head
	}

	for q.Next != nil {
		p, q = p.Next, q.Next
	}

	l := p.Next
	p.Next = l.Next
	l.Next = nil

	return head
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n1.Print()
	h := removeNthFromEnd(n1, 2)
	h.Print()
}
