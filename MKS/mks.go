// Merge k Sorted Lists
package MKS

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(a *ListNode, b *ListNode) *ListNode {
	var (
		head  *ListNode
		first *ListNode
		slave *ListNode
	)
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a.Val <= b.Val {
		first = a
		head = a
		slave = b
	} else {
		first = b
		head = b
		slave = a
	}

	for head.Next != nil && slave != nil {
		if slave.Val < head.Next.Val {
			head.Next, slave = slave, head.Next
		}
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	if slave != nil {
		head.Next = slave
	}
	return first
}

func listNodePrint(p *ListNode) {
	head := p
	for head != nil {
		fmt.Printf("-> %v ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func mergeKLists(lists []*ListNode) *ListNode {
	ll := len(lists)
	if ll == 0 {
		return nil
	} else if ll == 1 {
		return lists[0]
	} else if ll == 2 {
		return mergeTwoList(lists[0], lists[1])
	}
	ll = (ll + 1) / 2
	a := mergeKLists(lists[:ll])
	b := mergeKLists(lists[ll:])
	return mergeTwoList(a, b)
}

func MergeTest() {
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	b := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	c := ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	e := mergeKLists([]*ListNode{&a, &b, &c})
	listNodePrint(e)
}
