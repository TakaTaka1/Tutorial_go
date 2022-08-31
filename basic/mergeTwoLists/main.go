package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(li1 *ListNode, li2 *ListNode) *ListNode {
	dummyList := &ListNode{}

	var p = dummyList
	// [1,4,5] 1 < 2, 4 > 2, 4 > 3, 4 == 4
	// [2,3,4]
	for li1 != nil && li2 != nil {
		if li1.Val < li2.Val { // 小さい方が
			fmt.Println(p)
			p.Next = li1 // 現在のli1の値をp.Nextに上書きする
			fmt.Println(p.Next)
			li1 = li1.Next // li1を次のノードで上書きする
		} else {
			// fmt.Printf("li1.Val = %d : li2.Val = %d \n", li1.Val, li2.Val)
			p.Next = li2
			li2 = li2.Next
		}

		p = p.Next // li1かli2の値を格納した後にpのノードを進める
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
			Val: 4, Next: &ListNode{
				Val:  5,
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
		fmt.Println(test.Val)
		test = test.Next
	}

}
