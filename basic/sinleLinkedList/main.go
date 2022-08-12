package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type SingleLinkedList struct {
	len  int
	head *Node
}

func InitList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) Insert(val int) {
	node := &Node{value: val}
	if l.len != 0 {
		node.next = l.head
		l.head = node
		l.len++
	} else {
		l.head = node
		l.len++
	}
}

func (l *SingleLinkedList) InsertAt(val int, pos int) *Node {
	if l.head == nil {
		return nil
	}

	ptr := l.head
	node := &Node{value: val}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	node.next = ptr.next // 追加するnodeのアドレスに「元データの次の参照アドレス」を渡す
	ptr.next = node      // 追加するnodeのアドレスを渡す

	l.len++
	return ptr
}

func (l *SingleLinkedList) GetLinkedListDataAt(n int) *Node {
	if n > l.len {
		return nil
	}

	ptr := l.head
	for i := 0; i < n; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *SingleLinkedList) Traverse() []int {
	current := l.head
	if l.head == nil {
		fmt.Println("empty")
	}
	list := []int{}
	for current != nil {
		list = append(list, current.value)
		current = current.next
	}
	return list
}

func main() {
	list := InitList()
	list.Insert(10)
	list.Insert(11)
	list.Insert(12)

	list.InsertAt(20, 1)
	fmt.Println(list.Traverse())
}
