package BinarySearch

import "testing"

func TestMaxDistance(t *testing.T) {
	arr1 := []int{2}
	arr2 := []int{2, 2, 1}
	t.Log(maxDistance(arr1, arr2))
}
