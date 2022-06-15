package LinkedList

func partition(head *ListNode, x int) *ListNode {
	var dummySmallHead = &ListNode{}
	var dummyBigHead = &ListNode{}
	prevS := dummySmallHead
	prevB := dummyBigHead

	curr := head
	for curr != nil {
		if curr.Val < x {
			prevS.Next = curr
			prevS = curr
		} else {
			prevB.Next = curr
			prevB = curr
		}
		curr = curr.Next
	}
	prevB.Next = nil
	prevS.Next = dummyBigHead.Next

	return dummySmallHead.Next
}
