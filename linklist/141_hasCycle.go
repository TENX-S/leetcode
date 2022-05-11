package linklist

// T: O(n)
// S: O(1)
func hasCycle(head *ListNode) bool {
	// 注意这个开头的判空条件
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next // 注意fast的起始位置

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

// 推荐写这个，假如面到了这个题目
func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 使用哈希表
// T: O(n)
// S: O(n)
func hasCycle3(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}

	return false
}
