package linklist


func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    base := lists[0]
    for _, i := range lists[1:] {
        base = mergeTwoLists(base, i)
    }
    return base
}
