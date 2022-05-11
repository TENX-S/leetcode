package linklist

import "testing"

func TestMergeInBetween(t *testing.T) {
    res := mergeInBetween(NewFromList([]int{0, 1, 2, 3, 4, 5}), 3, 4, NewFromList([]int{1000000, 1000001, 1000002}))
    res.PrintList()
}
