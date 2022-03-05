package linklist

func middleNode(head *ListNode) *ListNode {
	fast := head
	for fast != nil && fast.Next != nil { // 注意判空条件 
		fast = fast.Next.Next
		head = head.Next
	}
	return head
}
