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

func printList(ls []*ListNode) {
	for i := 0; i < len(ls); i++ {
		fmt.Printf("%d:", i)
		ls[i].print()
	}
}
func min(x, y *ListNode) *ListNode {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	if x.Val <= y.Val {
		return x
	}
	return y
}
func FindMinNode(lists []*ListNode) (int, int) {
	index := 0
	minlistNode := lists[0]
	cnt := 0
	for i := 0; i < len(lists); i++ {
		minlistNode = min(minlistNode, lists[i])
		if minlistNode == lists[i] {
			index = i
		}
		if lists[i] != nil {
			cnt = cnt + 1
		}
	}
	return index, cnt
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var l0 *ListNode = nil
	var l *ListNode = nil
	idx, lens := FindMinNode(lists)
	if lens <= 1 {
		return lists[idx]
	}

	for ; lens >= 1; idx, lens = FindMinNode(lists) {

		if lens == 0 {
			break
		}
		if l0 == nil {
			l = lists[idx]
			l0 = l
		} else {
			fmt.Println(idx, lens)
			if lists[idx] != nil {
				l.Next = lists[idx]
				l = l.Next
			}
		}

		lists[idx] = lists[idx].Next

	}

	l.Next = lists[idx]
	return l0
}
func main() {
	/*
		l1 := &ListNode{Val: 1}
		l2 := &ListNode{Val: 2}
		l4 := &ListNode{Val: 4}
		l1.Next = l2
		l2.Next = l4

	*/

	/*
		l1 := &ListNode{Val: 1}
		l1.Next = &ListNode{Val: 4}
		l1.Next.Next = &ListNode{Val: 5}

		l2 := &ListNode{Val: 1}
		l2.Next = &ListNode{Val: 3}
		l2.Next.Next = &ListNode{Val: 4}

		l3 := &ListNode{Val: 2}
		l3.Next = &ListNode{Val: 6}
	*/
	lists := []*ListNode{nil, nil}

	r := mergeKLists(lists)
	r.print()

}
