package main

import "fmt"

// type SingleLinkedList struct {
// 	len  int
// 	head *ListNode
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(li1 *ListNode, li2 *ListNode) *ListNode {
	dummyList := &ListNode{}

	var p = dummyList

	for li1 != nil && li2 != nil {
		if li1.Val < li2.Val {
			p.Next = li1
			li1 = li1.Next
		} else {
			p.Next = li2
			li2 = li2.Next
		}
		p = p.Next
	}

	if li1 != nil {
		p.Next = li1
	} else {
		p.Next = li2
	}

	return dummyList.Next
}

func main() {

	li1 := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	li2 := &ListNode{
		Val: 2, Next: &ListNode{
			Val: 3, Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	test := mergeTwoLists(li1, li2)
	for test != nil {
		fmt.Println(test)
		test = test.Next
	}

}
