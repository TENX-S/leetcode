package linklist

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var prev *ListNode
	curr := slow.Next
	slow.Next = nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	p := head
	t := prev
	for prev != nil && p.Next != nil {
		next := t.Next
		t.Next = p.Next
		p.Next = t
		t = next
		p = p.Next.Next
	}
}
