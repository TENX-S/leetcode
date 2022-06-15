package BinarySearch

func findTheDistanceValue(arr1 []int, arr2 []int, d int) (res int) {
	quickSort(arr2)
	for _, i := range arr1 {
		low := i - d
		high := i + d
		if !exists(arr2, low, high) {
			res++
		}
	}

	return
}

func exists(arr []int, low, high int) bool {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= low && arr[mid] <= high {
			return true
		} else if arr[mid] < low {
			left = mid + 1
		} else if arr[mid] > high {
			right = mid
		}
	}

	return false
}

func quickSort(arr []int) {
	qSort(arr, 0, len(arr)-1)
}

func qSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
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
		if start < j {
			qSort(arr, start, j)
		}
		if end > i {
			qSort(arr, i, end)
		}
	}
}
