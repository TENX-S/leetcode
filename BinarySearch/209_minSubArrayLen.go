package BinarySearch

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	var res int
	var sum int
	for right < len(nums) {
		sum += nums[right]
		right++

		for sum > target {
			sum -= nums[left]
			left--
			res = min(res, left-right+1)
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
