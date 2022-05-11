package linklist

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    
    start, end := head, head

    for i := 0;i < k;i++ {
        if end == nil {
            return end
        }
        end = end.Next
    }

    newHead := reverseUntil(start, end)
    start.Next = reverseKGroup(end, k)

    return newHead
}

func reverseUntil(start, end *ListNode) *ListNode {
    var prev *ListNode
    curr := start 
    for curr != end {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    
    return prev
}
