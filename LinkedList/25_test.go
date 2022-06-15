package LinkedList

import "testing"

func TestReverseKGroup(t *testing.T) {
	reverseKGroup(NewFromList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 3).PrintList()
}
