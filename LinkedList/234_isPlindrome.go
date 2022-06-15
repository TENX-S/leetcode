package LinkedList

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	firstHalfEnd := firstHalf(head).Next
	secondHalfHead := reverseListInner(firstHalfEnd)

	x := head
	y := secondHalfHead

	for y != nil {
		if x.Val != y.Val {
			return false
		}
		x = x.Next
		y = y.Next
	}
	firstHalfEnd = reverseListInner(secondHalfHead)
	return true
}

func reverseListInner(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func firstHalf(head *ListNode) *ListNode {
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		head = head.Next
	}
	return head
}
