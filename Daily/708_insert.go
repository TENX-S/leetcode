package Daily

type Node struct {
	Val  int
	Next *Node
}

func CreateCircularLinkList(arr []int) *Node {
	dummyHead := Node{}
	p := &dummyHead
	for _, val := range arr {
		aNode := &Node{Val: val}
		p.Next = aNode
		p = p.Next
	}
	p.Next = dummyHead.Next

	return dummyHead.Next
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		var head Node
		head.Val = x
		head.Next = &head
		return &head
	}

	cand := &Node{Val: x}

	min := aNode
	p := aNode.Next
	for p != aNode {
		if p.Val < min.Val {
			min = p
		}
		p = p.Next
	}

	if min.Val > x {
		prev := aNode
		for prev.Next != min {
			prev = prev.Next
		}
		cand.Next = min
		prev.Next = cand

		return aNode
	}

	start := min

	for start.Next != min {
		if start.Next.Val >= x {
			cand.Next = start.Next
			start.Next = cand
			return aNode
		}
		start = start.Next
	}

	cand.Next = min
	start.Next = cand

	return aNode
}
