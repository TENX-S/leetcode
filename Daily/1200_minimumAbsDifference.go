package Daily

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var absDiff []int
	for i := 0; i+1 < len(arr); i++ {
		absDiff = append(absDiff, abs(arr[i+1]-arr[i]))
	}
	var ans [][]int

	var min int = math.MaxInt64
	for _, v := range absDiff {
		if v < min {
			min = v
		}
	}
	for i, v := range absDiff {
		if v == min {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
