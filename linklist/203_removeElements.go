package linklist

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	
	for tmp := dummyHead; tmp != nil; {
		 if tmp.Next.Val == val {
			 tmp.Next = tmp.Next.Next
		 } else {
			 tmp = tmp.Next
		 }
	}

	return dummyHead.Next
}