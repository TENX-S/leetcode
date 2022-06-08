package binarySearch

import "testing"

func TestKWeakestRows(t *testing.T) {
	//arr1 := []int{0, 0, 0, 0, 0}
	//arr2 := []int{1, 0, 0, 0, 0}
	//arr3 := []int{1, 1, 1, 1, 0}
	//arr4 := []int{1, 1, 1, 1, 1}
	//
	//t.Log(bSearchWithRightBound(arr1, 1))
	//t.Log(bSearchWithRightBound(arr2, 1))
	//t.Log(bSearchWithRightBound(arr3, 1))
	//t.Log(bSearchWithRightBound(arr4, 1))

	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}

	t.Log(kWeakestRows(mat, 3))
}
