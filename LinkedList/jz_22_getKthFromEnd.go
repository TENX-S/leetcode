package LinkedList

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head

	for k > 0 {
		if fast == nil {
			return nil
		}
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
