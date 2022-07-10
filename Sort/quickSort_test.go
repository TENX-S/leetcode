package Sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := [][]int{
		{},
		{-1},
		{1, 0},
		{10, 9, 8, 7, 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d\n", i), func(t *testing.T) {
			fmt.Println("Original: ", c)
			QuickSort(c)
			fmt.Println("Sorted: ", c)
			if !sort.IntsAreSorted(c) {
				t.Errorf("%d didn't pass", i)
			}
		})
	}
}
