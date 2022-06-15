package LinkedList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dummyHead := &Node{}
	p := dummyHead
	for head != nil {
		node := &Node{Val: head.Val}
		random := &Node{Val: head.Random.Val}
		p.Next = node
		p.Next.Random = random
		p = p.Next
		head = head.Next
	}
	return dummyHead.Next
}

func copyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	p := dummyHead
	for head != nil {
		node := &ListNode{Val: head.Val}
		p.Next = node
		p = p.Next
		head = head.Next
	}
	return dummyHead.Next
}
