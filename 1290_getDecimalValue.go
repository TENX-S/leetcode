package main

func getDecimalValue(head *ListNode) (ans int) {
	for head != nil {
		ans = (ans << 1) + head.Val
		head = head.Next
	}
	return
}

