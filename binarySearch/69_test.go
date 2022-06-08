package binarySearch

import "testing"

func TestMySqrt(t *testing.T) {

	arr := createNumbers(0, 16)
	for i := range arr {
		t.Logf("%d -> %d\n", i, mySqrt(i))
	}
}

func createNumbers(lo int, hi int) []int {
	s := make([]int, hi-lo+1)

	for i := range s {
		s[i] = i + lo
	}

	return s
}
