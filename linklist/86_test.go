package linklist

import "testing"

func TestPartition(t *testing.T) {
    l := NewFromList([]int{1, 4, 3, 2, 2, 5})
    res := partition(l, 3)
    res.PrintList()
}
