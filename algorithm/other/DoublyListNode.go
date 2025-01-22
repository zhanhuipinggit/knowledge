package main

import "fmt"

type DoublyListNode struct {
	pre, next *DoublyListNode
	value     int
}

func NewDoublyListNode(value int) *DoublyListNode {
	return &DoublyListNode{value: value}
}

func CreateDoublyLinkedList(arr []int) *DoublyListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := NewDoublyListNode(arr[0])
	current := head
	for i := 1; i < len(arr); i++ {
		node := NewDoublyListNode(arr[i])
		current.next = node
		node.pre = current
		current = current.next
	}
	return head
}

func main() {
	arr := []int{1, 2, 3, 4}
	head := CreateDoublyLinkedList(arr)
	// 在双链表头部插入新节点 0
	nodeInsert := NewDoublyListNode(0)
	nodeInsert.next = head
	head.pre = nodeInsert
	head = nodeInsert

	// 在双链表尾部插入新元素6
	nodeInsertEnd := NewDoublyListNode(5)
	tail := head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = nodeInsertEnd
	nodeInsertEnd.pre = tail
	// 更新尾节点引用
	tail = nodeInsertEnd

	// 在第 3 个节点后面插入新节点 66
	// 找到第 3 个节点
	p := head
	for i := 0; i < 2; i++ {
		p = p.next
	}
	nodeInsertMiddle := NewDoublyListNode(66)
	next := p.next
	p.next = nodeInsertMiddle
	nodeInsertMiddle.pre = p
	nodeInsertMiddle.next = next
	next.pre = nodeInsertMiddle

	pr := head
	for pr.next != nil {
		pr = pr.next
		fmt.Println("v", pr.value)
	}

	// 删除第3个节点
	dp := head
	for i := 0; i < 2; i++ {
		dp = dp.next
	}

	dp.next = dp.next.next
	dp.next.pre = dp

	// 双向链表删除第一个元素
	dfirst := head.next
	dfirst.pre = nil
	head.next = nil
	head = dfirst

}
