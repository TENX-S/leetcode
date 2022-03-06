package linklist

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// 获取链表长度
func (head *ListNode) GetLength() int {
	p := head
	var length int
	for p != nil {
		length++
		p = p.Next
	}

	return length
}

// 遍历打印链表
func (head *ListNode) PrintList() {
	p := head

	for p != nil {
		fmt.Printf("%d\t", p.Val)
		p = p.Next
	}
}

// 打印某个节点
func (head *ListNode) PrintNode() {
	if head != nil {
		fmt.Printf("Nil node")
	} else {
		fmt.Printf("%+v", head)
	}
}

// 删除指定位置的节点
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

// 在指定位置后插入数据
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

