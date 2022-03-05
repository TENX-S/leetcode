package linklist


// T: O(m+n)
// S: O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyhead := &ListNode{}
	p := dummyhead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	// 特别注意这个边界条件！
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return dummyhead.Next
}

