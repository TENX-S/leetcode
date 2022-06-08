package binarySearch

import "testing"

func TestSearch(t *testing.T) {
	t.Log(search([]int{4, 5, 6, 7, 0, 1, 2}, 7))
}
