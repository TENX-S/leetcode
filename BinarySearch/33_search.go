package BinarySearch

func search(nums []int, target int) int {
	peek := findPeek(nums)
	if target < nums[0] {
		res := bSearch(nums[peek+1:], target)
		if res != -1 {
			return res + peek + 1
		} else {
			return -1
		}
	} else {
		return bSearch(nums[0:peek+1], target)
	}
}

func findPeek(arr []int) int {
	left, right, pivot := 0, len(arr), arr[0]

	for left < right {
		mid := left + (right-left)/2
		key := arr[mid]
		if key > pivot {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left != 0 {
		return left - 1
	}
	return left
}
