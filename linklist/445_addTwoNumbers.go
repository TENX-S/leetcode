package linklist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    r1, r2 := reverseList(l1), reverseList(l2)
    dummyHead := &ListNode{}
    curr := dummyHead
    carry := 0

    for r1 != nil || r2 != nil {
        node := &ListNode{Val: carry}

        if r1 != nil {
            node.Val += r1.Val
            r1 = r1.Next
        }

        if r2 != nil {
            node.Val += r2.Val
            r2 = r2.Next
        }

        carry = node.Val / 10
        node.Val %= 10
        curr.Next = node
        curr = curr.Next
    }

    if carry != 0 {
        curr.Next = &ListNode{Val: carry}
    }

    return reverseList(dummyHead.Next)
}
