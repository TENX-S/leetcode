package LinkedList

// 双指针法
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	p := prev.Next

	for p != nil {
		if p.Val == val {
			prev.Next = p.Next
		}
		prev = prev.Next
		p = p.Next
	}

	return dummyHead.Next
}

// 单指针法
func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val { // 解决第一个节点的问题
		return head.Next
	}

	p := head
	for p.Next != nil && p.Next.Val != val {
		p = p.Next
	}

	p.Next = p.Next.Next

	return head
}
