package BinarySearch

func specialArray(nums []int) int {
	quickSort(nums)
	left, right := 0, len(nums)-1
	var mid int
	var prev int
	for left <= right {
		mid = left + (right-left)/2
		key := nums[mid]
		x := len(nums) - mid
		if key >= x {
			prev = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	x := len(nums) - prev
	if nums[prev] >= x {
		if bSearch(nums[:prev], x) != -1 {
			return -1
		}
		return x
	}

	return -1
}
