package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type SingleLinkedList struct {
	len        int
	head, tail *Node
}

func InitList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) InsertHead(val int) {
	node := &Node{value: val}
	if l.len != 0 {
		node.next = l.head // 新規ノードがheadを指すように
		l.head = node      // headがnodeを指すように
		l.len++
	} else {
		l.head = node
		l.len++
	}
}

func (l *SingleLinkedList) InsertTail(val int) {
	node := &Node{value: val}
	if l.len != 0 {
		l.tail.next = node // tailの次ノードを新規追加のnodeにする
		l.len++
	} else {
		l.head = node
		l.len++
	}
	l.tail = node // tailを現在のノードにする(上書きされるわけではなく新たなノード(tail)が作成されるイメージ)
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
	list.InsertTail(10)
	list.InsertTail(11)
	list.InsertTail(12)
	// list.InsertAt(20, 1)
	fmt.Println(list.Traverse())
}
