package binarySearch

func maxDistance(nums1, nums2 []int) (ans int) {
	if nums1[len(nums1)-1] > nums2[0] {
		return
	}

	for i, n1 := range nums1 {
		j := lowerBound(nums2, n1)
		if j == -1 {
			continue
		}
		ans = max(ans, j-i)
	}

	return
}

func lowerBound(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		key := arr[mid]
		if target <= key {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == 0 || arr[left-1] < target {
		return -1
	}
	return left - 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
