package LinkedList

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // 注意这里是:=
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev // 注意返回值
}
