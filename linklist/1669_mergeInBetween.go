package linklist

import "fmt"

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    l := list1.GetLength()
    aLeft := getOffsetLeft(list1, l-a)
    bNode := getOffsetLeft(list1, l-b-1)
    fmt.Println(bNode)
    aLeft.Next = list2
    rest := bNode.Next
    bNode.Next = nil

    p := list1
    for p.Next != nil {
        p = p.Next
    }
    p.Next = rest

    return list1
}

func getOffsetLeft(head *ListNode, offset int) *ListNode {
    slow, fast := head, head
    offset++
    for offset > 0 {
        fast = fast.Next
        offset--
    }

    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    return slow
}
