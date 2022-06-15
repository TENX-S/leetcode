package LinkedList

import "testing"

func TestCopyLinkList(t *testing.T) {
	l := NewFromList([]int{1, 2, 3, 4, 5, 6})
	n := copyList(l)
	n.PrintList()
}
