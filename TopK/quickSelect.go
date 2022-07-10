package TopK

func findKthLargest(nums []int, k int) int {
	return Partition(nums, k)
}

func Partition(arr []int, k int) int { //快排
	i, j := 0, len(arr)-1
	key := arr[0]
	for i <= j {
		for arr[i] < key {
			i++
		}
		for arr[j] > key {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	if i+1 == k {
		return key
	} else if i+1 < k {
		return Partition(arr[i+1:], k-i-1)
	} else {
		return Partition(arr[:i], k)
	}
}
