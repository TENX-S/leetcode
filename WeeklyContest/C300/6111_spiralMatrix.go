package C300

import (
	. "github.com/TENX-S/findwork/LinkedList"
)

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
		for k := range arr[i] {
			arr[i][k] = -1
		}
	}
	upper, lower := 0, m-1
	left, right := 0, n-1

	for head != nil {
		if upper <= lower {
			for j := left; j <= right && head != nil; j++ {
				arr[upper][j] = head.Val
				head = head.Next
			}
			upper++
		}
		if left <= right {
			for i := upper; i <= lower && head != nil; i++ {
				arr[i][right] = head.Val
				head = head.Next
			}
			right--
		}
		if upper <= lower {
			for j := right; j >= left && head != nil; j-- {
				arr[lower][j] = head.Val
				head = head.Next
			}
			lower--
		}
		if left <= right {
			for i := lower; i >= upper && head != nil; i-- {
				arr[i][left] = head.Val
				head = head.Next
			}
			left++
		}
	}
	return arr
}
