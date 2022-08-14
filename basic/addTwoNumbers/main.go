package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	elem := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + elem
		newNode := ListNode{val % 10, nil}
		elem = val / 10
		// val<10の場合0になる
		// 20>val>10の場合1になる

		if head == nil || tail == nil {
			head = &newNode
			tail = &newNode
		} else {
			tail.Next = &newNode
			tail = tail.Next
		}
		l1 = l1.Next
		l2 = l2.Next

	}

	for l1 != nil {
		val := l1.Val + elem
		fmt.Println(val)
		newNode := ListNode{val % 10, nil}
		elem = val / 10
		if head == nil {
			head = &newNode
			tail = &newNode
		} else {
			tail.Next = &newNode
			tail = tail.Next
		}
		l1 = l1.Next

	}
	for l2 != nil {
		val := l2.Val + elem
		newNode := ListNode{val % 10, nil}
		elem = val / 10

		if head == nil {
			head = &newNode
			tail = &newNode
		} else {
			tail.Next = &newNode
			tail = tail.Next
		}
		l2 = l2.Next
	}

	// 桁溢れがあれば最後に追加する
	if elem > 0 {
		newNode := ListNode{elem, nil}
		if head == nil {
			head = &newNode
			tail = &newNode
		} else {
			tail.Next = &newNode
			tail = tail.Next
		}
	}

	return head
}
