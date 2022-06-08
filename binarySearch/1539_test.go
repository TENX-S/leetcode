package binarySearch

import "testing"

func TestFindKthPositive(t *testing.T) {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	t.Log(findKthPositive(arr, k))
}
