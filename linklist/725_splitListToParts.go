package linklist

func splitListToParts(head *ListNode, k int) []*ListNode {
	var res []*ListNode
	l := head.GetLength()
	curr := head
	if l <= k {
		for curr != nil {
			next := curr.Next
			curr.Next = nil
			res = append(res, curr)
			curr = next
		}

		for i := k - l; i > 0; i-- {
			res = append(res, (*ListNode)(nil))
		}
		return res
	}

	for i := 0; i < l%k; i++ {
		res = append(res, curr)
		for j := 0; j < l/k; j++ {
			curr = curr.Next
		}
		next := curr.Next
		curr.Next = nil
		curr = next
	}

	for i := 0; i < k-l%k; i++ {
		res = append(res, curr)
		for j := 0; j < l/k-1; j++ {
			curr = curr.Next
		}
		next := curr.Next
		curr.Next = nil
		curr = next
	}

	return res
}
