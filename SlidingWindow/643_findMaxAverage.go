package SlidingWindow

import "math"

type NumArr struct {
	preSum []int
}

func NewNumArr(nums []int) NumArr {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArr{preSum: preSum}
}

func (n NumArr) sumRange(left, right int) float64 {
	return float64(n.preSum[right+1] - n.preSum[left])
}

func findMaxAverage(nums []int, k int) (ans float64) {
	max := float64(math.MinInt64)
	preArr := NewNumArr(nums)
	for i := 0; i <= len(nums)-k; i++ {
		cand := preArr.sumRange(i, i+k-1) / float64(k)
		if cand > max {
			max = cand
		}
	}
	return max
}
