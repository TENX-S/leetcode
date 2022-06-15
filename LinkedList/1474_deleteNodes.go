package LinkedList

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	p := head
	for {
		for i := 0; i < m-1 && p != nil; i++ {
			p = p.Next
		}
		for i := 0; i < n && p != nil && p.Next != nil; i++ {
			p.Next = p.Next.Next
		}
		if p == nil {
			break
		}
		p = p.Next
	}
	return head
}

// 这个更好
func deleteNodes2(head *ListNode, m int, n int) *ListNode {
	pre, last := head, head

	for pre != nil && last != nil {
		for i := 0; i < m-1; i++ {
			if pre == nil {
				break
			}
			pre = pre.Next
		}

		for i := 0; i < m+n; i++ {
			if last == nil {
				break
			}
			last = last.Next
		}

		if pre != nil {
			pre.Next = last
			pre = last
		}
	}

	return head
}
