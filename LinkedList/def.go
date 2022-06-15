package LinkedList

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// GetLength 获取链表长度
func (head *ListNode) GetLength() int {
	p := head
	var length int
	for p != nil {
		length++
		p = p.Next
	}

	return length
}

// PrintList 遍历打印链表
func (head *ListNode) PrintList() {
	p := head

	for p.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	fmt.Println(p.Val)
}

// PrintNode 打印某个节点
func (head *ListNode) PrintNode() {
	if head != nil {
		fmt.Printf("Nil node")
	} else {
		fmt.Printf("%+v", head)
	}
}

// Delete 删除指定位置的节点
func (head *ListNode) Delete(index int) (success bool) {
	if index < 0 || index > head.GetLength() {
		fmt.Println("Invalid index!")
	} else {
		p := head
		for i := 0; i < index; i++ {
			p = p.Next
		}
		p.Next = p.Next.Next
		success = true
	}
	return
}

// Insert 在指定位置后插入数据
func (head *ListNode) Insert(index int, data int) (success bool) {
	if index < 0 || index > head.GetLength() {
		fmt.Println("Invalid data")
	} else {
		p := head
		for i := 0; i < index; i++ {
			p = p.Next
		}
		var NewNode ListNode
		NewNode.Val = data
		NewNode.Next = p.Next
		p.Next = &NewNode
		success = true
	}
	return
}

func NewFromList(arr []int) *ListNode {
	dummyHead := &ListNode{Next: nil}
	p := dummyHead
	for _, i := range arr {
		node := ListNode{
			Val:  i,
			Next: nil,
		}
		p.Next = &node
		p = p.Next
	}
	return dummyHead.Next
}
